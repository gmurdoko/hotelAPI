package controllers

import (
	"encoding/json"
	"hotelAPI/main/masters/usecases"
	"hotelAPI/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RoomHandler app
type RoomHandler struct {
	RoomUsecase usecases.RoomUsecase
}

//RoomController app
func RoomController(r *mux.Router, s usecases.RoomUsecase) {
	roomHandler := RoomHandler{s}
	cat := r.PathPrefix("/rooms").Subrouter()
	cat.HandleFunc("", roomHandler.ListRooms).Methods(http.MethodGet)
	cat.HandleFunc("/available", roomHandler.ListAvailableRooms).Methods(http.MethodGet)
	cat.HandleFunc("/booked", roomHandler.ListBookedRooms).Methods(http.MethodGet)

}

//ListRooms app
func (s *RoomHandler) ListRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := s.RoomUsecase.GetAllRoom()
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		roomResponse.Status = http.StatusNotFound
		roomResponse.Message = "Not Found"
		roomResponse.Data = err
		log.Println(err)
		w.Write([]byte("Data Not Found"))
	} else {
		roomResponse.Status = http.StatusOK
		roomResponse.Message = "Get All Room Success"
		roomResponse.Data = rooms

		byteOfRooms, err := json.Marshal(roomResponse)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Opps, Something Wrong"))
		}
		w.Write([]byte(byteOfRooms))
	}
	log.Println("Endpoint hit: Get All Room")
}

//ListAvailableRooms app
func (s *RoomHandler) ListAvailableRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := s.RoomUsecase.GetAvailableRoom()
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		roomResponse.Status = http.StatusNotFound
		roomResponse.Message = "Not Found"
		roomResponse.Data = err
		log.Println(err)
		w.Write([]byte("Data Not Found"))
	} else {
		roomResponse.Status = http.StatusOK
		roomResponse.Message = "Get Available Room Success"
		roomResponse.Data = rooms

		byteOfRooms, err := json.Marshal(roomResponse)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Opps, Something Wrong"))
		}
		w.Write([]byte(byteOfRooms))
	}
	log.Println("Endpoint hit: Get Available Rooms")
}

//ListBookedRooms app
func (s *RoomHandler) ListBookedRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := s.RoomUsecase.GetBookedRoom()
	var roomResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		roomResponse.Status = http.StatusNotFound
		roomResponse.Message = "Not Found"
		roomResponse.Data = err
		log.Println(err)
		w.Write([]byte("Data Not Found"))
	} else {
		roomResponse.Status = http.StatusOK
		roomResponse.Message = "Get Booked Room Success"
		roomResponse.Data = rooms

		byteOfRooms, err := json.Marshal(roomResponse)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Opps, Something Wrong"))
		}
		w.Write([]byte(byteOfRooms))
	}
	log.Println("Endpoint hit: Get Booked Rooms")
}
