package models

type Room struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	IDGame   int    `json:"id_game"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Participant struct {
	ID        int `json:"id"`
	IDRoom    int `json:"id_room"`
	IDAccount int `json:"id_account"`
}

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type RoomParticipant struct {
	Room        Room        `json:"room"`
	Participant Participant `json:"participant"`
	Account     Account     `json:"account"`
}

type RoomParticipants struct {
	Room         Room          `json:"room"`
	Participants []Participant `json:"participants"`
	Account      Account       `json:"account"`
}

type RoomParticipantsResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []RoomParticipant `json:"data"`
}

// type ParticipantResponse struct {
// 	Status  int    `json:"status"`
// 	Message string `json:"message"`
// 	Data    Room   `json:"data"`
// }

type ParticipantsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}
