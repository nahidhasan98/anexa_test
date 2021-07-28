package main

import (
	"anexa_test/api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//using Gorilla mux router
	r := mux.NewRouter()

	//a message for ensuring that local server is running
	fmt.Println("Local Server is running on port 8001 ...")

	//creating api endpoints
	r.HandleFunc("/api/cart", api.GetCartInfo).Methods("GET")                //GET
	r.HandleFunc("/api/cart", api.AddItemToCart).Methods("POST")             //POST
	r.HandleFunc("/api/cart/{id}", api.DeleteItemFromCart).Methods("DELETE") //DELETE
	r.HandleFunc("/api/cart", api.ResetCart).Methods("DELETE")               //DELETE
	r.HandleFunc("/api/order", api.PlaceOrder).Methods("POST")               //POST

	//for localhost server
	http.ListenAndServe(":8001", r)
}
