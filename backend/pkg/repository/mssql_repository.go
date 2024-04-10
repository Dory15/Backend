package repository

import (
	"backend/pkg/models"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

var dbConnectionString string

type MSSQL struct {
}

func NewMSSQLRepository() IRepository {

	return MSSQL{}
}

func (mssql MSSQL) openDatabaseConnection() (*sql.DB, error) {
	LoadEnv()
	return sql.Open("mssql", dbConnectionString)
}

func (mssql MSSQL) closeDatabaseConnection(databaseConnection *sql.DB) error {
	return databaseConnection.Close()
}

func (mssql MSSQL) SaveUser(user models.SignUpStruct) error {

	databaseConnection, err := mssql.openDatabaseConnection()
	if err != nil {
		return errors.New("No se pudo crear la conexion a MSSQL " + err.Error())
	}
	defer mssql.closeDatabaseConnection(databaseConnection)

	_, err = databaseConnection.Query("EXEC insert_user @email = ?, @user = ?, @password = ?, @phone = ?",
		user.Email, user.User, user.Password, user.Phone)
	if err != nil {
		return errors.New("No se pudo ejecutar el store procedure insert_user: " + err.Error())

	}

	return nil
}

func (mssql MSSQL) GetUser(user models.SignUpStruct) (models.SignUpStruct, error) {

	databaseConnection, err := mssql.openDatabaseConnection()
	if err != nil {
		return models.SignUpStruct{}, errors.New("No se pudo crear la conexion a MSSQL " + err.Error())
	}
	defer mssql.closeDatabaseConnection(databaseConnection)

	dbResult, err := databaseConnection.Query("SELECT TOP 1 email, user, password, phone FROM [user] WHERE email = ? OR phone = ? OR user = ?",
		user.Email, user.Phone, user.Email)
	if err != nil {
		return models.SignUpStruct{}, errors.New("No se pudo ejecutar el store procedure insert_user: " + err.Error())

	}

	var dbUser models.SignUpStruct

	if dbResult.Next() {
		dbResult.Scan(&dbUser.Email, &dbUser.User, &dbUser.Password, &dbUser.Phone)
	}

	return dbUser, nil
}

func LoadEnv() {
	dbConnectionString = os.Getenv("DB_STRING")
	fmt.Println(dbConnectionString)
	if dbConnectionString == "" {
		panic("No hay cadena de conexion a base de datos")
	}
}
