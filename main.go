package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bmv126/go-app/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// GET
	router.HandleFunc("/member/{name}", api.GetMember).Methods("GET")
	router.HandleFunc("/allMembers", api.GetAllMembers).Methods("GET")

	// POST
	router.HandleFunc("/addMember", api.AddMember).Methods("POST")
	router.HandleFunc("/addMembers", api.AddMembers).Methods("POST")

	// DELETE
	router.HandleFunc("/deleteMember/{name}", api.DeleteMember).Methods("DELETE")
	router.HandleFunc("/deleteAll", api.DeleteAll).Methods("DELETE")

	// PUT
	router.HandleFunc("/modifyMember/{name}", api.ModifyMember).Methods("PUT")

	fmt.Println("Starting server....")
	log.Fatal(http.ListenAndServe(":5000", router))

}
