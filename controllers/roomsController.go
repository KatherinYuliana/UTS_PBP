package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "UTS/models"

	"github.com/gorilla/mux"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM rooms"

	// Read from Query Param
	roomName := r.URL.Query()["room_name"]
	idGame := r.URL.Query()["id_game"]
	if roomName != nil {
		fmt.Println(roomName[0])
		query += "WHERE room_name='" + roomName[0] + "'"
	}

	if idGame != nil {
		if roomName[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " id_game='" + idGame[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		sendErrorResponse(w)
		return
	}
	var room m.Room
	var rooms []m.Room
	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.RoomName, &room.IDGame); err != nil {
			sendErrorResponse(w)
			return
		} else {
			rooms = append(rooms, room)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// 2 queries
	roomRows, err := db.Query("SELECT * FROM rooms")
	if err != nil {
		sendErrorResponse(w)
		return
	}
	var room m.Room
	var rooms []m.Room
	for roomRows.Next() {
		if err := roomRows.Scan(&room.ID, &room.RoomName, &room.IDGame); err != nil {
			sendErrorResponse(w)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	participantRows, err := db.Query("SELECT * FROM participants")
	if err != nil {
		sendErrorResponse(w)
		return
	}
	var participant m.RoomParticipant
	var participants []m.RoomParticipant
	for participantRows.Next() {
		if err := participantRows.Scan(&participant.Room.ID, &participant.Room.RoomName, &participant.Participant.ID, &participant.Participant.IDAccount); err != nil {
			sendErrorResponse(w)
			return
		} else {
			participants = append(participants, participant)
		}
	}

	// 1 query with join
	query := `SELECT r.id, r.room_name, a.username, p.id, p.id_account FROM rooms r, accounts a JOIN participants p ON r.id = p.id_room AND a.id = p.id_account`
	roomParticipantRow, err := db.Query(query)
	if err != nil {
		sendErrorResponse(w)
		return
	}
	var roomParticipant m.RoomParticipant
	var roomParticipants []m.RoomParticipant
	for roomParticipantRow.Next() {
		if err := roomParticipantRow.Scan(
			&roomParticipant.Room.ID, &roomParticipant.Room.RoomName, &roomParticipant.Participant.ID,
			&roomParticipant.Participant.IDAccount, &roomParticipant.Account.Username); err != nil {
			print(err.Error())
			sendErrorResponse(w)
			return
		} else {
			roomParticipants = append(roomParticipants, roomParticipant)
		}
	}

	var response m.RoomParticipantsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = roomParticipants
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		return
	}
	roomName := r.Form.Get("room_name")
	idGame, _ := strconv.Atoi(r.Form.Get("id_game"))

	_, errQuery := db.Exec("INSERT INTO rooms(room_name, id_game) values (?,?)",
		roomName,
		idGame,
	)

	var response m.RoomResponse
	if errQuery == nil {
		sendSuccessResponse(w)
	} else {
		sendErrorResponse(w)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	participantId := vars["participant_id"]

	_, errQuery := db.Exec("DELETE FROM participants WHERE id=?",
		participantId,
	)

	if errQuery == nil {
		sendSuccessResponse(w)
	} else {
		sendErrorResponse(w)
	}
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response m.ErrorResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
