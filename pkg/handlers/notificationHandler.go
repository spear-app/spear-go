package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/spear-app/spear-go/pkg/domain/notification"
	errs "github.com/spear-app/spear-go/pkg/err"
	"github.com/spear-app/spear-go/pkg/service"
)

type NotificationHandlers struct {
	service service.NotificationService
}
var (
	validate *validator.Validate
)


func (notificationHandler NotificationHandlers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var notiObj *notification.Notification
	json.NewDecoder(r.Body).Decode(&notiObj)

	validate = validator.New()
	fmt.Println(notiObj.Title, notiObj.Body, notiObj.UserUID)
	err := validate.Struct(notiObj)

	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}
	
	err = notificationHandler.service.Create(notiObj)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notiObj)
}


func (notificationHandler NotificationHandlers) ReadByNotificationID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicatsion/json")

	params := mux.Vars(r) // Get params
	strId := params["id"]
	id, err := strconv.Atoi(strId)

	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse("invalid notification id", http.StatusInternalServerError))
		return
	}

	var notiObj notification.Notification
	notiObj, err = notificationHandler.service.ReadByNotificationID(id)
	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errs.NewResponse("notification not found", http.StatusNotFound))
		return
	case nil:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(notiObj)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		return
	}
}