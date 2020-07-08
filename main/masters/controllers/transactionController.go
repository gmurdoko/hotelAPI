package controllers

import (
	"encoding/json"
	"hotelAPI/config"
	"hotelAPI/main/masters/models"
	"hotelAPI/main/masters/usecases"
	"hotelAPI/main/middleware/token"
	"hotelAPI/utils"
	"log"
	"net/http"

	// "github.com/gmurdoko/middleware/middleware"
	"github.com/gorilla/mux"
)

// ReserveHandler app
type ReserveHandler struct {
	ReserveUsecase usecases.ReserveUsecase
}

//ReserveController app
func ReserveController(r *mux.Router, s usecases.ReserveUsecase) {
	reserveHandler := ReserveHandler{s}

	rss := r.PathPrefix("/reserves").Subrouter()
	rve := r.PathPrefix("/reserve").Subrouter()
	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		rss.Use(token.TokenValidationMiddleware)
		rve.Use(token.TokenValidationMiddleware)
		detailReserveController(rss, rve, reserveHandler)
	} else {
		detailReserveController(rss, rve, reserveHandler)
	}

}

func detailReserveController(rss, rve *mux.Router, reserveHandler ReserveHandler) {
	rss.HandleFunc("", reserveHandler.ListReserves).Methods(http.MethodGet)
	rve.HandleFunc("", reserveHandler.PostReserve).Methods(http.MethodPost)
}

//ListReserves app
func (s *ReserveHandler) ListReserves(w http.ResponseWriter, r *http.Request) {
	reserves, err := s.ReserveUsecase.GetAllReserve()
	var reserveResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		reserveResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", Data: err.Error()}
		utils.ResponseWrite(&reserveResponse, w)
		log.Println(err)
	} else {
		reserveResponse = utils.Response{Status: http.StatusOK, Message: "Get All Room Success", Data: reserves}
		utils.ResponseWrite(&reserveResponse, w)
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
		reserveResponse = utils.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		utils.ResponseWrite(&reserveResponse, w)
		log.Println(err)
	} else {
		reserveResponse = utils.Response{Status: http.StatusOK, Message: "Get All Room Success", Data: inReserve}
		utils.ResponseWrite(&reserveResponse, w)
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
