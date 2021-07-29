package api

import (
	"anexa_test/db"
	"anexa_test/model"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PlaceOrder func will insert an order to the DB and return the order id
func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	//declaring variable in which data will be stored
	var orderData model.Order

	//connecting to DB
	DB, ctx, cancel := db.Connect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	//selecting DB collection/table & taking to a variable
	cartColl := DB.Collection("cart")
	orderColl := DB.Collection("order")

	//first retrieving all item(s) from the cart from DB
	cursor, err := cartColl.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	//getting multiple documents(rows)
	//Iterating through the cursor allows us to decode one document at a time
	for cursor.Next(ctx) {
		//creating a temporary variable in which the single document can be decoded
		var temp model.Item
		err := cursor.Decode(&temp)
		if err != nil {
			fmt.Println(err)
		}
		orderData.Total += temp.Total //price of each item will be added together and stored to totalPrice variable

		orderData.Items = append(orderData.Items, temp) //finally taking this single item to the slice of items
	}

	//got item(s) from cart, now checking if the cart is epty or not

	//preparing data for json response
	var returnData model.ResponseData

	if orderData.Items == nil { //if the cart is empty
		returnData.Status = "error"
		returnData.Message = "No product found in the cart. Cart is empty."
	} else { //got some item(s) in the cart
		//now inserting order (document/row) to DB
		res, err := orderColl.InsertOne(ctx, orderData)
		if err != nil {
			fmt.Println(err)
		}

		//getting inserted id as string from the result of insert query to DB
		insertedID := res.InsertedID.(primitive.ObjectID).Hex() //type assertion && Calling Hex func

		returnData.Status = "success"
		returnData.ID = insertedID
		returnData.Message = "Order successfully placed. Order id: " + insertedID

		//now resetting the cart after making an order
		//deleting all item(s) (documents/rows) from cart
		_, err = cartColl.DeleteMany(ctx, bson.M{})
		if err != nil {
			fmt.Println(err)
		}
	}

	//encoding data to json
	rData, err := json.Marshal(returnData)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //setting content type as application/json
	w.Write(rData)                                     //finally response back to client
}
