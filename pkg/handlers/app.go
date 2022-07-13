package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/spear-app/spear-go/pkg/domain/authen"
	"github.com/spear-app/spear-go/pkg/domain/notification"
	"github.com/spear-app/spear-go/pkg/driver"
	"github.com/spear-app/spear-go/pkg/middleware"
	"github.com/spear-app/spear-go/pkg/service"
	"github.com/spear-app/spear-go/pkg/utils"
)

func Start() {

	router := mux.NewRouter()
	//this CORS to enable frontend request to the backend endpoints
	headers := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	dbConnection := driver.GetDbConnetion()
	driver.Seed(dbConnection)

	authenHandler := AuthenHandlers{service.NewAuthenService(authen.NewAuthenRepositoryDb(dbConnection))}
	notificationHandler := NotificationHandlers{service.NewNotificationService(notification.NewNotificationRepositoryDb(dbConnection))}
	//authorization endpoints
	router.HandleFunc("/api/signup", authenHandler.Signup).Methods(http.MethodPost)
	router.HandleFunc("/api/login", authenHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/profile/{id:[0-9]+}", middleware.TokenVerifyMiddleware(authenHandler.ReadUser)).Methods(http.MethodGet)
	router.HandleFunc("/api/auth/profile/{id:[0-9]+}", middleware.TokenVerifyMiddleware(authenHandler.Update)).Methods(http.MethodPut)
	router.HandleFunc("/api/auth/profile/{id:[0-9]+}", middleware.TokenVerifyMiddleware(authenHandler.Delete)).Methods(http.MethodDelete)
	router.Path("/api/v1/confirmEmail/{id}").HandlerFunc(authenHandler.VerifyEmail).Methods(http.MethodPost)
	router.HandleFunc("/api/notification/create", middleware.TokenVerifyMiddleware(notificationHandler.Create)).Methods(http.MethodPost)
	router.HandleFunc("/api/notification/getNotificationById/{id:[0-9]+}", middleware.TokenVerifyMiddleware(notificationHandler.ReadByNotificationID)).Methods(http.MethodGet)
	router.HandleFunc("/api/notification/getNotificationByUserId/{id:[0-9]+}", middleware.TokenVerifyMiddleware(notificationHandler.ReadByUserID)).Methods(http.MethodGet)
	router.HandleFunc("/api/audio/send_audio", middleware.TokenVerifyMiddleware(Wav)).Methods(http.MethodPost)
	router.HandleFunc("/api/audio/start_conversation", middleware.TokenVerifyMiddleware(StartConversation)).Methods(http.MethodPost)
	router.HandleFunc("/api/audio/end_conversation", middleware.TokenVerifyMiddleware(EndConversation)).Methods(http.MethodPost)
	router.HandleFunc("/api/audio/recorded_audio", middleware.TokenVerifyMiddleware(RecordedAudio)).Methods(http.MethodPost)

	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("00:00").Do(func() { utils.DeleteJobInternal(dbConnection) })

	s.StartAsync()
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))
}
