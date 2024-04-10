package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CarlosMore29/env_cm"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	// "creditassigner/pkg/controllers"
	"backend/pkg/controllers"
	"backend/pkg/logger"
)

var PORT string

var fileName string = "main"
var mainLogger logger.ILogger

func init() {
	envGlobals()
	mainLogger = logger.NewLoggerInstace(fileName)
}

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.SetHeader("Content-Type", "application/json"), // Set Content-Type Header to application/json
		middleware.Logger,          // Log API request calls
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		//middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Route("/", func(r chi.Router) {
		r.Mount("/", controllers.Routes())
	})

	return router
}

func main() {

	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		mainLogger.Info(fmt.Sprintf("%s %s\n", method, route))
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		mainLogger.Error(err.Error())
	}

	mainLogger.Error(http.ListenAndServe(":"+PORT, router).Error())

}

func envGlobals() {
	env_cm.GetEnvFile()
	PORT = os.Getenv("PORT")
}
