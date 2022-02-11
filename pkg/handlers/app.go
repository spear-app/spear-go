package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/spear-app/spear-go/pkg/domain/authen"
	"github.com/spear-app/spear-go/pkg/driver"
	"github.com/spear-app/spear-go/pkg/middleware"
	"github.com/spear-app/spear-go/pkg/service"
)

func Start() {
	router := mux.NewRouter()
	//this CORS to enable frontend request to the backend endpoints
	headers := handlers.AllowedHeaders([]string{"Content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	dbConnection := driver.GetDbConnetion()
	driver.Seed(dbConnection)

	authenHandler := AuthenHandlers{service.NewAuthenService(authen.NewAuthenRepositoryDb(dbConnection))}
	//authorization endpoints
	router.HandleFunc("/api/signup", authenHandler.Signup).Methods(http.MethodPost)
	router.HandleFunc("/api/login", authenHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/profile/{id:[0-9]+}", middleware.TokenVerifyMiddleware(authenHandler.ReadUser)).Methods(http.MethodGet)
	router.HandleFunc("/api/auth/profile/{id:[0-9]+}", middleware.TokenVerifyMiddleware(authenHandler.Update)).Methods(http.MethodPut)
	router.HandleFunc("/api/auth/profile/{id:[0-9]+}", middleware.TokenVerifyMiddleware(authenHandler.Delete)).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", handlers.CORS(headers, methods, origins)(router)))

}
