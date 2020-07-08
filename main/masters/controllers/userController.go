package controllers

import (
	"encoding/json"
	"hotelAPI/main/masters/models"
	"hotelAPI/main/masters/usecases"
	"hotelAPI/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// UserHandler app
type UserHandler struct {
	userUsecase usecases.UserUsecase
}

//UserController app
func UserController(r *mux.Router, s usecases.UserUsecase) {
	userHandler := UserHandler{s}
	aut := r.PathPrefix("/auth").Subrouter()
	aut.HandleFunc("", userHandler.AuthUser).Methods(http.MethodPost)
	usr := r.PathPrefix("/user").Subrouter()
	usr.HandleFunc("", userHandler.PostUser).Methods(http.MethodPost)

}

//AuthUser app
func (u *UserHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var data models.Users
	_ = json.NewDecoder(r.Body).Decode(&data)
	isValid, _ := u.userUsecase.GetAuthUsers(&data)

	if isValid {
		w.WriteHeader(http.StatusOK)
		token, err := utils.JwtEncoder(data.UserName, "rahasiadong")
		if err != nil {
			http.Error(w, "Failed token generation", http.StatusUnauthorized)
		} else {
			w.Write([]byte(token))
		}
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}
}

//PostUser app
func (u *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var inUser models.Users
	var userResponse utils.Response
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&inUser)
	if err != nil {
		log.Println(err)
		w.Write([]byte("cant read JSON"))
	}
	err = u.userUsecase.PostUser(&inUser)
	if err != nil {
		userResponse = utils.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		utils.ResponseWrite(&userResponse, w)
		log.Println(err)
	} else {
		userResponse = utils.Response{Status: http.StatusAccepted, Message: "Post User Success", Data: inUser}
		utils.ResponseWrite(&userResponse, w)
	}
	log.Println("Endpoint hit: Post User")
}
