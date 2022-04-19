package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/spear-app/spear-go/pkg/domain/notification"
	errs "github.com/spear-app/spear-go/pkg/err"
	"github.com/spear-app/spear-go/pkg/service"
)

type NotificationHandlers struct {
	service service.NotificationService
}


func (notificationHandler NotificationHandlers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var notiObj *notification.Notification
	json.NewDecoder(r.Body).Decode(&notiObj)
	var userIdStr string = strconv.FormatUint(uint64(notiObj.UserUID), 10)
	_, err := strconv.Atoi(userIdStr)
	if err!=nil{
		//TODO response with bad request 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse("invalid user id", http.StatusBadRequest))
		return
	}
	
	err = notificationHandler.service.Create(notiObj)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusInternalServerError))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notiObj)
}