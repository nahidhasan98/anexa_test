package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Product for holding a single product details
type Product struct {
	ID          string   `bson:"_id,omitempty"`
	Code        string   `bson:"code"`
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	Price       float64  `bson:"price"`
	Count       int      `bson:"count"`
	Discount    float64  `bson:"discount"`
	Colors      []string `bson:"colors"`
	Sizes       []string `bson:"sizes"`
}

//Item for holding a single item details
type Item struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Code  string             `json:"code" bson:"code"`
	Unit  int                `json:"unit" bson:"unit"`
	Price float64            `json:"price" bson:"price"`
	Total float64            `json:"total" bson:"total"`
}

//Order for holding a single order details
type Order struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Items []Item             `json:"items" bson:"items"`
	Total float64            `json:"total" bson:"total"`
}

//ResponseData for holding a api response details
type ResponseData struct {
	Status  string `json:"status"`
	ID      string `json:"id,omitempty"`
	Message string `json:"message"`
}
