package masters

import (
	"database/sql"
	"hotelAPI/main/masters/controllers"
	"hotelAPI/main/masters/repositories"
	"hotelAPI/main/masters/usecases"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB) {
	//Rooms
	roomRepo := repositories.InitRoomRepoImpl(db)
	roomUsecase := usecases.InitRoomUsecaseImpl(roomRepo)
	controllers.RoomController(r, roomUsecase)

	//Transaction
	reserveRepo := repositories.InitReserveRepoImpl(db)
	reserveUsecase := usecases.InitReserveUsecaseImpl(reserveRepo)
	controllers.ReserveController(r, reserveUsecase)
}
