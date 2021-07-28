package api

import (
	"anexa_test/db"
	"anexa_test/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PlaceOrder func will insert an order to the DB and return the order id
func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	//declaring variable in which data will be stored
	var orderData model.Order

	//receiving request body
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(rBody, &orderData) //getting json data to the variable

	//calculating total price of each item based on price and unit
	//and total oreder price
	for key, val := range orderData.Items {
		orderData.Items[key].Total = val.Price * float64(val.Unit)
		orderData.Total += orderData.Items[key].Total
	}

	//connecting to DB
	DB, ctx, cancel := db.Connect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	//selecting DB collection/table & taking to a variable
	order := DB.Collection("order")

	//inserting order (document/row) to DB
	res, err := order.InsertOne(ctx, orderData)
	if err != nil {
		fmt.Println(err)
	}

	//getting inserted id as string from the result of insert query to DB
	insertedID := res.InsertedID.(primitive.ObjectID).Hex() //type assertion && Calling Hex func

	//preparing data for json response
	returnData := struct {
		Status  string `json:"status"`
		ID      string `json:"id"`
		Message string `json:"message"`
	}{
		Status:  "success",
		ID:      insertedID,
		Message: "Order sucseccfully placed. Order id: " + insertedID,
	}

	//encoding data to json
	rData, err := json.Marshal(returnData)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //setting content type as application/json
	w.Write(rData)                                     //finally response back to client
}
