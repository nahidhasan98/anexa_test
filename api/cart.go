package api

import (
	"anexa_test/db"
	"anexa_test/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetCartInfo func will return all item of the cart with total price
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	//declaring variable in which data will be stored
	var cartData []model.Item
	var totalPrice float64

	//connecting to DB
	DB, ctx, cancel := db.Connect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	//selecting DB collection/table & taking to a variable
	cartColl := DB.Collection("cart")

	//finding all documents/rows from DB
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
		totalPrice += temp.Total //price of each item will be added together and stored to totalPrice variable

		cartData = append(cartData, temp) //finally taking this single document to the slice of documents
	}

	//preparing data for json response
	returnData := struct {
		Items []model.Item `json:"items"`
		Total float64      `json:"total"`
	}{
		Items: cartData,
		Total: totalPrice,
	}

	//encoding data to json
	rData, err := json.Marshal(returnData)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //setting content type as application/json
	w.Write(rData)                                     //finally response back to client
}

//AddItemToCart func will insert an item to the cart and return the inserted id
func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	//declaring variable in which data will be stored
	var item model.Item

	//receiving request body
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(rBody, &item) //getting json data to the variable

	//calculating total price of this item based on price and unit
	item.Total = item.Price * float64(item.Unit)

	//connecting to DB
	DB, ctx, cancel := db.Connect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	//selecting DB collection/table & taking to a variable
	cartColl := DB.Collection("cart")

	//inserting item (document/row) to DB
	res, err := cartColl.InsertOne(ctx, item)
	if err != nil {
		fmt.Println(err)
	}

	//getting inserted id as string from the result of insert query to DB
	insertedID := res.InsertedID.(primitive.ObjectID).Hex() //type assertion && Calling Hex func

	//preparing data for json response
	returnData := model.ResponseData{
		Status:  "success",
		ID:      insertedID,
		Message: "Item successfully added to cart. Inserted id: " + insertedID,
	}

	//encoding data to json
	rData, err := json.Marshal(returnData)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //setting content type as application/json
	w.Write(rData)                                     //finally response back to client
}

//DeleteItemFromCart func will delete an item from the cart and return the number of deleted item with the id of deleted item
func DeleteItemFromCart(w http.ResponseWriter, r *http.Request) {
	//getting item id from request parameter
	params := mux.Vars(r)
	id := params["id"]

	//converting requested id from string to primitive.ObjectID
	idPrim, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}

	//connecting to DB
	DB, ctx, cancel := db.Connect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	//selecting DB collection/table & taking to a variable
	cartColl := DB.Collection("cart")

	//deleting specific item (document/row) from DB
	res, err := cartColl.DeleteOne(ctx, bson.M{"_id": idPrim})
	if err != nil {
		fmt.Println(err)
	}

	//preparing data for json response
	var returnData model.ResponseData

	if res.DeletedCount == 0 { //if no item found with the provided id
		returnData.Status = "error"
		returnData.Message = "No product found with ID: " + id
	} else { //if item found
		returnData.Status = "success"
		returnData.ID = id
		returnData.Message = "Item successfully deleted from cart. Total item deleted: " + fmt.Sprintf("%d", res.DeletedCount)
	}

	//encoding data to json
	rData, err := json.Marshal(returnData)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //setting content type as application/json
	w.Write(rData)                                     //finally response back to client
}

//ResetCart func will delete all item(s) from the cart and return the number of deleted item(s)
func ResetCart(w http.ResponseWriter, r *http.Request) {
	//connecting to DB
	DB, ctx, cancel := db.Connect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	//selecting DB collection/table & taking to a variable
	cartColl := DB.Collection("cart")

	//deleting all item(s) (documents/rows) from cart
	res, err := cartColl.DeleteMany(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	//preparing data for json response
	var returnData model.ResponseData

	if res.DeletedCount == 0 { //if cart is empty
		returnData.Status = "error"
		returnData.Message = "No product found. Cart is already empty."
	} else { //got some item(s) in the cart and deleted
		returnData.Status = "success"
		returnData.Message = "All item(s) successfully deleted from cart. Total item(s) deleted: " + fmt.Sprintf("%d", res.DeletedCount)
	}

	//encoding data to json
	rData, err := json.Marshal(returnData)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //setting content type as application/json
	w.Write(rData)                                     //finally response back to client
}
