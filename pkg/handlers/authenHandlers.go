package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/spear-app/spear-go/pkg/domain/user"
	emailVerification "github.com/spear-app/spear-go/pkg/emailVerfication"
	errs "github.com/spear-app/spear-go/pkg/err"
	"github.com/spear-app/spear-go/pkg/service"
	"github.com/spear-app/spear-go/pkg/utils"
	"log"
	"net/http"
	"net/mail"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"

	"github.com/spear-app/spear-go/pkg/domain/authen"
)

type AuthenHandlers struct {
	service service.AuthenService
}

//utility function to check if the email valid
func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

//utility function to validate the email and password
func validateEmailAndPassword(userObj user.User) error {
	//err error()
	if userObj.Email == "" {
		return errs.ErrEmailMissing
	}
	if !valid(userObj.Email) {
		return errs.ErrInvalidEmail
	}
	if userObj.Password == "" {
		return errs.ErrMissingPassword
	}
	return nil
}

func (authenHandler AuthenHandlers) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//extracting usr obj
	var userObj user.User
	json.NewDecoder(r.Body).Decode(&userObj)
	//validating email and password
	err := validateEmailAndPassword(userObj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}
	//validating name and gender
	err = validateNameAndGender(userObj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}
	//encrypting password
	hash, err := bcrypt.GenerateFromPassword([]byte(userObj.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		return
	}
	userObj.Password = string(hash)
	//database connection
	err = authenHandler.service.Signup(&userObj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		return
	}

	//generating email confirmation code
	code := emailVerification.CodeGenerator()

	//hashing email confirmation code
	hashOTP, err := bcrypt.GenerateFromPassword([]byte(code), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		return
	}

	//inserting the hashed code into the database
	userObj.OTP = string(hashOTP)
	err = receiver.service.InsertOTP(userObj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	// sending the code to the user
	err = emailVerification.SendEmail(userObj.Email, code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	//generating the token
	token, err := utils.GenerateToken(userObj)
	if err != nil {
		log.Fatal(err)
	}
	type Data struct {
		Token string    `json:"token"`
		User  user.User `json:"user"`
	}
	var data Data
	userObj.Password = ""
	userObj.OTP = ""
	data.User = userObj
	data.Token = token
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (authenHandler AuthenHandlers) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var auth authen.Authen
	json.NewDecoder(r.Body).Decode(&auth.User)
	//validate email and password
	err := validateEmailAndPassword(auth.User)
	//handling errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}
	password := auth.User.Password
	err = authenHandler.service.Login(&auth.User)
	//handling errors
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errs.NewResponse("No record found", http.StatusNotFound))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//check if password is entered password matching with actually user password
	if !utils.CheckPasswordHash(password, auth.User.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrInvalidPassword.Error(), http.StatusUnauthorized))
		return
	}

	//generating the token
	token, err := utils.GenerateToken(auth.User)
	if err != nil {
		log.Fatal(err)
	}
	type Data struct {
		Token string    `json:"token"`
		User  user.User `json:"user"`
	}
	var data Data
	auth.User.Password = ""
	data.User = auth.User

	data.Token = token
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (authenHandler AuthenHandlers) ReadUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Init user
	var authen authen.Authen
	id := params["id"]
	tempId, _ := strconv.Atoi(id)
	authen.User.ID = uint(tempId)
	err := authenHandler.service.ReadUserByID(&authen.User)
	//handling errors
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoRowsFound.Error(), http.StatusNotFound))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&authen.User)
}

func (authenHandler AuthenHandlers) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var usrObj user.User
	_ = json.NewDecoder(r.Body).Decode(&usrObj)
	// validate inputs
	err := validateNameAndGender(usrObj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusBadRequest))
		return
	}
	idStr, _ := strconv.Atoi(id)
	usrObj.ID = uint(idStr)
	err = authenHandler.service.Update(&usrObj)

	//handling errors
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoRowsFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrServerErr.Error(), http.StatusInternalServerError))
		}
		return
	}
	usrObj.Password = ""
	usrObj.Email = ""

	type Data struct {
		Message string    `json:"message"`
		User    user.User `json:"user"`
	}
	var data Data
	data.Message = "update done successfully"
	data.User = usrObj
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// Delete endpoint to delete user
func (authenHandler AuthenHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	err := authenHandler.service.Delete(id)
	//handling errors
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoRowsFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errs.NewResponse("User has been deleted successfully", http.StatusOK))
}

func validateNameAndGender(userObj user.User) error {
	//err error()
	if userObj.Name == "" {
		return errs.ErrMissingName
	}
	if userObj.Gender == "" {
		return errs.ErrMissingGender
	}
	return nil
}
