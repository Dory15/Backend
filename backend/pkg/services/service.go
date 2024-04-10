package services

import (
	"backend/pkg/models"
	repositoryLib "backend/pkg/repository"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/CarlosMore29/env_cm"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	Email string
	jwt.RegisteredClaims
}

var jwtClaims jwt.RegisteredClaims
var secret string

type Service struct {
	repository repositoryLib.IRepository
}

func init() {
	LoadEnv()
}

func NewService() Service {
	return Service{repository: repositoryLib.NewRepository()}
}

func (service Service) SignUp(singUpStruct models.SignUpStruct) error {

	hashedPassword, err := hashPassword(singUpStruct.Password)
	if err != nil {
		return err
	}

	singUpStruct.Password = hashedPassword

	dbUser, err := service.repository.GetUser(singUpStruct)

	if err != nil {
		return err
	}

	if dbUser.Email != "" {
		return errors.New("El correo/telefono ya se encuentra registrado")
	}

	err = service.repository.SaveUser(singUpStruct)

	if err != nil {
		return err
	}

	return nil
}

func (service Service) SignIn(singUpStruct models.SignUpStruct) (string, error) {

	dbUser, err := service.repository.GetUser(singUpStruct)

	if err != nil {
		return "", err
	}

	if dbUser.Email == "" {
		return "", errors.New("usuario / contraseña incorrectos")
	}

	match := checkPasswordHash(singUpStruct.Password, dbUser.Password)

	if !match {
		return "", errors.New("usuario / contraseña incorrectos")
	}

	return createJwt(dbUser.Email)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createJwt(email string) (string, error) {
	claims := CustomClaims{
		email,
		jwtClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func LoadEnv() {
	env_cm.GetEnvFile()
	secret = os.Getenv("JWT_SECRET")
	expirationString := os.Getenv("JWT_EXPIRATION")

	expirationNumberInSeconds, err := strconv.Atoi(expirationString)
	if err != nil {
		panic(err)
	}

	jwtClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(int64(expirationNumberInSeconds))))

}
