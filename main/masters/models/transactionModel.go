package models

// Reserves app
type Reserves struct {
	Nik          string `json:"nik"`
	CustomerName string `json:"customerName"`
	RoomID       string `json:"roomID"`
	RoomName     string `json:"roomName"`
	BookedAt     string `json:"bookedAt"`
	EndedAt      string `json:"endedAt"`
}
