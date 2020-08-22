package models

//Rooms app
type Rooms struct {
	ID        int    `json:"id"`
	RoomName  string `json:"roomName"`
	Price     string `json:"price"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
