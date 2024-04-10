package controllers

import (
	"backend/pkg/dtos"
	"backend/pkg/logger"
	"backend/pkg/models"
	"backend/pkg/services"
	"backend/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var fileName string = "controller"
var controllerLogger logger.ILogger

var service services.Service

func init() {
	controllerLogger = logger.NewLoggerInstace(fileName)
	service = services.NewService()
}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/signup", signUp)
	router.Post("/signin", signIn)
	return router
}

func signUp(res http.ResponseWriter, req *http.Request) {

	var registroStruct dtos.RegistroStruct

	res.Header().Set("Content-Type", "application/json")
	jsonResponseWriter := json.NewEncoder(res)

	err := json.NewDecoder(req.Body).Decode(&registroStruct)

	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = registroStruct.Validate()
	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		errorResponse := dtos.ErrorStruct{Error: err.Error()}
		jsonResponseWriter.Encode(errorResponse)
		return
	}

	signUpStruct := models.SignUpStruct{
		Email:    registroStruct.Correo,
		User:     registroStruct.Usuario,
		Password: registroStruct.Contrasena,
		Phone:    registroStruct.Telefono,
	}

	err = service.SignUp(signUpStruct)

	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		errorResponse := dtos.ErrorStruct{Error: err.Error()}
		jsonResponseWriter.Encode(errorResponse)
		return
	}

	res.WriteHeader(http.StatusCreated)
}

func signIn(res http.ResponseWriter, req *http.Request) {

	var loginStruct dtos.LoginStruct

	res.Header().Set("Content-Type", "application/json")
	jsonResponseWriter := json.NewEncoder(res)

	err := json.NewDecoder(req.Body).Decode(&loginStruct)

	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = loginStruct.Validate()
	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		errorResponse := dtos.LoginResponse{Error: err.Error()}
		jsonResponseWriter.Encode(errorResponse)
		return
	}

	signUpStruct := models.SignUpStruct{
		Password: loginStruct.Contrasena,
	}

	err = utils.ValidateEmail(loginStruct.UsuarioOCorreo)
	if err != nil {
		signUpStruct.User = loginStruct.UsuarioOCorreo
	} else {
		signUpStruct.Email = loginStruct.UsuarioOCorreo
	}

	token, err := service.SignIn(signUpStruct)

	if err != nil {
		controllerLogger.Error(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		errorResponse := dtos.LoginResponse{Error: err.Error()}
		jsonResponseWriter.Encode(errorResponse)
		return
	}

	res.WriteHeader(http.StatusAccepted)
	loginResponse := dtos.LoginResponse{Token: token}
	jsonResponseWriter.Encode(loginResponse)
}
