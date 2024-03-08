package main

import (
	"UTS/controllers"
	"fmt"
	"log"
	"net/http"

	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	// rooms
	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/rooms/{id_room}", controllers.GetDetailRooms).Methods("GET")
	router.HandleFunc("/rooms", controllers.InsertRoom).Methods("POST")
	router.HandleFunc("/participants/{participant_id}", controllers.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
