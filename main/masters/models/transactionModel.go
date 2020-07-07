package models

// Reserves app
type Reserves struct {
	Nik          string `json:"nik"`
	CustomerName string `json:"customerName"`
	RoomID       int    `json:"roomID"`
	RoomName     string `json:"roomName"`
	BookedAt     string `json:"bookedAt"`
	EndedAt      string `json:"endedAt"`
}
