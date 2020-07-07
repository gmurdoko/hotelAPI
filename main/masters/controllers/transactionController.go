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

// ReserveHandler app
type ReserveHandler struct {
	ReserveUsecase usecases.ReserveUsecase
}

//ReserveController app
func ReserveController(r *mux.Router, s usecases.ReserveUsecase) {
	reserveHandler := ReserveHandler{s}
	cat := r.PathPrefix("/reserves").Subrouter()
	cat.HandleFunc("", reserveHandler.ListReserves).Methods(http.MethodGet)
	// cat.HandleFunc("/available", roomHandler.ListAvailableRooms).Methods(http.MethodGet)
	// cat.HandleFunc("/booked", roomHandler.ListBookedRooms).Methods(http.MethodGet)
	cat = r.PathPrefix("/reserve").Subrouter()
	cat.HandleFunc("", reserveHandler.PostReserve).Methods(http.MethodPost)
	// cat.HandleFunc("", roomHandler.PutRoom).Methods(http.MethodPut)
	// cat.HandleFunc("/{id}", roomHandler.DeleteRoom).Methods(http.MethodDelete)

}

//ListReserves app
func (s *ReserveHandler) ListReserves(w http.ResponseWriter, r *http.Request) {
	reserves, err := s.ReserveUsecase.GetAllReserve()
	var reserveResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		reserveResponse.Status = http.StatusNotFound
		reserveResponse.Message = "Not Found"
		reserveResponse.Data = err
		log.Println(err)
		w.Write([]byte("Data Not Found"))
	} else {
		reserveResponse.Status = http.StatusOK
		reserveResponse.Message = "Get All Reserve Success"
		reserveResponse.Data = reserves

		byteOfReserves, err := json.Marshal(reserveResponse)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Opps, Something Wrong"))
		}
		w.Write([]byte(byteOfReserves))
	}
	log.Println("Endpoint hit: Get All Reserve")
}

// //ListAvailableRooms app
// func (s *RoomHandler) ListAvailableRooms(w http.ResponseWriter, r *http.Request) {
// 	rooms, err := s.RoomUsecase.GetAvailableRoom()
// 	var roomResponse utils.Response
// 	w.Header().Set("content-type", "application/json")
// 	if err != nil {
// 		roomResponse.Status = http.StatusNotFound
// 		roomResponse.Message = "Not Found"
// 		roomResponse.Data = err
// 		log.Println(err)
// 		w.Write([]byte("Data Not Found"))
// 	} else {
// 		roomResponse.Status = http.StatusOK
// 		roomResponse.Message = "Get Available Room Success"
// 		roomResponse.Data = rooms

// 		byteOfRooms, err := json.Marshal(roomResponse)
// 		if err != nil {
// 			log.Println(err)
// 			w.Write([]byte("Opps, Something Wrong"))
// 		}
// 		w.Write([]byte(byteOfRooms))
// 	}
// 	log.Println("Endpoint hit: Get Available Rooms")
// }

// //ListBookedRooms app
// func (s *RoomHandler) ListBookedRooms(w http.ResponseWriter, r *http.Request) {
// 	rooms, err := s.RoomUsecase.GetBookedRoom()
// 	var roomResponse utils.Response
// 	w.Header().Set("content-type", "application/json")
// 	if err != nil {
// 		roomResponse.Status = http.StatusNotFound
// 		roomResponse.Message = "Not Found"
// 		roomResponse.Data = err
// 		log.Println(err)
// 		w.Write([]byte("Data Not Found"))
// 	} else {
// 		roomResponse.Status = http.StatusOK
// 		roomResponse.Message = "Get Booked Room Success"
// 		roomResponse.Data = rooms

// 		byteOfRooms, err := json.Marshal(roomResponse)
// 		if err != nil {
// 			log.Println(err)
// 			w.Write([]byte("Opps, Something Wrong"))
// 		}
// 		w.Write([]byte(byteOfRooms))
// 	}
// 	log.Println("Endpoint hit: Get Booked Rooms")
// }

//PostReserve app
func (s *ReserveHandler) PostReserve(w http.ResponseWriter, r *http.Request) {
	var inReserve models.Reserves
	var reserveResponse utils.Response
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&inReserve)
	if err != nil {
		log.Println(err)
		w.Write([]byte("cant read JSON"))
	}
	err = s.ReserveUsecase.PostReserve(&inReserve)
	if err != nil {
		reserveResponse.Status = http.StatusNotFound
		reserveResponse.Message = "Not Found"
		reserveResponse.Data = err
		log.Println(err)
		w.Write([]byte("Data Not Found"))
	} else {
		reserveResponse.Status = http.StatusOK
		reserveResponse.Message = "Post Reserve Success"
		reserveResponse.Data = inReserve
		byteOfReserves, err := json.Marshal(reserveResponse)
		if err != nil {
			w.Write([]byte("Opps, Something Wrong"))
		}
		w.Write([]byte(byteOfReserves))
	}

	log.Println("Endpoint hit: Post Reserve")
}

// //PutRoom app
// func (s *RoomHandler) PutRoom(w http.ResponseWriter, r *http.Request) {
// 	var inRoom models.Rooms
// 	var roomResponse utils.Response
// 	w.Header().Set("content-type", "application/json")
// 	err := json.NewDecoder(r.Body).Decode(&inRoom)
// 	if err != nil {
// 		w.Write([]byte("cant read JSON"))
// 	}
// 	err = s.RoomUsecase.PutRoom(&inRoom)
// 	if err != nil {
// 		roomResponse.Status = http.StatusNotFound
// 		roomResponse.Message = "Not Found"
// 		roomResponse.Data = err
// 		log.Println(err)
// 		w.Write([]byte("Data Not Found"))
// 	} else {
// 		roomResponse.Status = http.StatusOK
// 		roomResponse.Message = "Put Room Success"
// 		roomResponse.Data = inRoom
// 		byteOfRooms, err := json.Marshal(roomResponse)
// 		if err != nil {
// 			log.Println(err)
// 			w.Write([]byte("Opps, Something Wrong"))
// 		}
// 		w.Write([]byte(byteOfRooms))
// 	}
// 	log.Println("Endpoint hit: Put Room")
// }

// // DeleteRoom app
// func (s *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
// 	ex := mux.Vars(r)
// 	idINT, err := strconv.Atoi(ex["id"])
// 	var roomResponse utils.Response
// 	w.Header().Set("content-type", "application/json")
// 	err = s.RoomUsecase.DeleteRoom(idINT)
// 	if err != nil {
// 		roomResponse.Status = http.StatusNotFound
// 		roomResponse.Message = "Not Found"
// 		roomResponse.Data = err
// 		log.Println(err)
// 		w.Write([]byte("Data Not Found"))
// 	} else {
// 		roomResponse.Status = http.StatusOK
// 		roomResponse.Message = "Delete Room Success"
// 		roomResponse.Data = idINT
// 		byteOfCategories, err := json.Marshal(roomResponse)
// 		if err != nil {
// 			log.Println(err)
// 			w.Write([]byte("Opps, Something Wrong"))
// 		}
// 		w.Write([]byte(byteOfCategories))
// 	}

// 	log.Println("Endpoint hit: Delete room")
// }
