package controllers

import (
	"encoding/json"
	"fmt"
	"hotelAPI/config"
	"hotelAPI/main/masters/models"
	"hotelAPI/main/masters/usecases"
	"hotelAPI/main/middleware/token"
	"hotelAPI/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// RoomHandler app
type RoomHandler struct {
	RoomUsecase usecases.RoomUsecase
}

//RoomController app
func RoomController(r *mux.Router, s usecases.RoomUsecase) {
	roomHandler := RoomHandler{s}
	rooms := r.PathPrefix("/rooms").Subrouter()
	room := r.PathPrefix("/room").Subrouter()
	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		rooms.Use(token.TokenValidationMiddleware)
		room.Use(token.TokenValidationMiddleware)
		detailRoomController(rooms, room, roomHandler)
	} else {
		detailRoomController(rooms, room, roomHandler)
	}
}

func detailRoomController(rooms, room *mux.Router, roomHandler RoomHandler) {
	//Jamak
	rooms.HandleFunc("", roomHandler.ListRooms).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}", "status", "{status}", "orderBy", "{orderBy}", "sort", "{sort}").Methods(http.MethodGet)
	rooms.HandleFunc("/available", roomHandler.ListAvailableRooms).Methods(http.MethodGet)
	rooms.HandleFunc("/booked", roomHandler.ListBookedRooms).Methods(http.MethodGet)
	//Satuan
	room.HandleFunc("", roomHandler.PostRoom).Methods(http.MethodPost)
	room.HandleFunc("", roomHandler.PutRoom).Methods(http.MethodPut)
	room.HandleFunc("/{id}", roomHandler.DeleteRoom).Methods(http.MethodDelete)
}

//ListRooms app
func (s *RoomHandler) ListRooms(w http.ResponseWriter, r *http.Request) {
	offset := mux.Vars(r)["page"]
	limit := mux.Vars(r)["limit"]
	status := mux.Vars(r)["status"]
	orderBy := mux.Vars(r)["orderBy"]
	sort := mux.Vars(r)["sort"]
	keyword := mux.Vars(r)["keyword"]
	fmt.Println("keyword:", keyword, "offset:", offset, "limit:", limit, "status:", status, "orderBy:", orderBy, "sort:", sort)
	rooms, totalField, err := s.RoomUsecase.GetAllRoom(keyword, offset, limit, status, orderBy, sort)
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		roomResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", TotalField: *totalField, Data: err.Error()}
		utils.ResponseWrite(&roomResponse, w)
		log.Println(err)
	} else {
		roomResponse = utils.Response{Status: http.StatusOK, Message: "Get All Room Success", TotalField: *totalField, Data: rooms}
		utils.ResponseWrite(&roomResponse, w)
	}
	log.Println("Endpoint hit: Get All Room")
}

//ListAvailableRooms app
func (s *RoomHandler) ListAvailableRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := s.RoomUsecase.GetAvailableRoom()
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		roomResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", Data: err.Error()}
		utils.ResponseWrite(&roomResponse, w)
		log.Println(err)
	} else {
		roomResponse = utils.Response{Status: http.StatusOK, Message: "Get Available Room Success", Data: rooms}
		utils.ResponseWrite(&roomResponse, w)
	}
	log.Println("Endpoint hit: Get Available Rooms")
}

//ListBookedRooms app
func (s *RoomHandler) ListBookedRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := s.RoomUsecase.GetBookedRoom()
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		roomResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", Data: err.Error()}
		utils.ResponseWrite(&roomResponse, w)
		log.Println(err)
	} else {
		roomResponse = utils.Response{Status: http.StatusOK, Message: "Get Booked Room Success", Data: rooms}
		utils.ResponseWrite(&roomResponse, w)
	}
	log.Println("Endpoint hit: Get Booked Rooms")
}

//PostRoom app
func (s *RoomHandler) PostRoom(w http.ResponseWriter, r *http.Request) {
	var inRoom models.Rooms
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&inRoom)
	if err != nil {
		log.Println(err)
		// w.Write([]byte("cant read JSON"))
	}
	err = s.RoomUsecase.PostRoom(&inRoom)
	if err != nil {
		roomResponse = utils.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		utils.ResponseWrite(&roomResponse, w)
		log.Println(err)
	} else {
		roomResponse = utils.Response{Status: http.StatusAccepted, Message: "Post Room Success", Data: inRoom}
		utils.ResponseWrite(&roomResponse, w)
	}
	log.Println("Endpoint hit: Post Room")
}

//PutRoom app
func (s *RoomHandler) PutRoom(w http.ResponseWriter, r *http.Request) {
	var inRoom models.Rooms
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&inRoom)
	if err != nil {
		// w.Write([]byte("cant read JSON"))
	}
	err = s.RoomUsecase.PutRoom(&inRoom)
	if err != nil {
		roomResponse = utils.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		utils.ResponseWrite(&roomResponse, w)
		log.Println(err)
	} else {
		roomResponse = utils.Response{Status: http.StatusAccepted, Message: "Put Room Success", Data: inRoom}
		utils.ResponseWrite(&roomResponse, w)
	}
	log.Println("Endpoint hit: Put Room")
}

// DeleteRoom app
func (s *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	ex := mux.Vars(r)
	idINT, err := strconv.Atoi(ex["id"])
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	err = s.RoomUsecase.DeleteRoom(idINT)
	if err != nil {
		roomResponse = utils.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		utils.ResponseWrite(&roomResponse, w)
		log.Println(err)
	} else {
		roomResponse = utils.Response{Status: http.StatusAccepted, Message: "Delete Room Success", Data: idINT}
		utils.ResponseWrite(&roomResponse, w)
	}

	log.Println("Endpoint hit: Delete room")
}
