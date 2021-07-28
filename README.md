# Shopping Cart

A shopping cart for a fictional e-commerce site. Main objective is to create REST APIs for some specific request/endpoints.

I am using **Go** as the programming language. For database, I am using **MongoDB** (MongoDB Atlas).

**Models** used in app:

```go
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
```


## Build & Run

1. Clone the repository (inside the `src` directory of your `GOPATH`) using following command:

```bash
git clone https://github.com/nahidhasan98/anexa_test.git
```
2. Go inside the project directory:

```bash
cd anexa_test
```

3. Build the project. This will create an executable file named `app.exe` inside the current directory.

```bash
go build -o app.exe
```

4. And finally run the executable file.

```bash
app.exe [for cmd]
./app.exe [for Windows PowerShell]
```

## Usage

### 1. Add an item to cart

To add an item to the cart, an HTTP **POST** request should be made at `/api/cart` this endpoint. And a `json` data (showing below) should be included with the request body.

```json
{
	"code": "BDS03101",
	"unit": 2,
	"price": 202.8
}
```

Total price for this item in the cart will be calculated `(total = price * unit)` in backend thus no need to to provide `total` as `json` field.

This api call returns a `json` data as response in the following format:

```json
{
    "status": "success",
    "id": "<a hex encoding objectID>",
    "message": "Item successfully added to cart. Inserted id: <a hex encoding objectID>"
}
```

### 2. Delete an item from cart

To delete an item from the cart, an HTTP **DELETE** request should be made at `/api/cart/{id}` this endpoint where id should be an item id in the cart. The reason for providing an item id is to identify which item should be deleted from the cart.

This api call returns a `json` data as response in the following format:

```json
{
    "status": "success",
    "id": "6101c76e63bba269e95e1c5c",
    "message": "Item successfully deleted from cart. Total item deleted: 1"
}
```

If the specific item is not found in the cart that is requested to be deleted then the response is like:

```json
{
    "status": "error",
    "message": "No product found with ID: 6101c76e63bba269e95e1c5c"
}
```

### 3. Get cart info (All item(s) of the cart and total price)

To get cart information, an HTTP **GET** request should be made at `/api/cart` this endpoint.

This api call returns a `json` data as response in the following format:

```json
{
    "items": [
        {
            "id": "6101c76e63bba269e95e1c5c",
            "code": "CAF05179",
            "unit": 2,
            "price": 115.8,
            "total": 231.6
        },
        {
            "id": "6101c833277da1da855d4c76",
            "code": "BDS03101",
            "unit": 1,
            "price": 202,
            "total": 202
        }
    ],
    "total": 433.6
}
```

If the cart is empty then the response is like:

```json
{
    "items": null,
    "total": 0
}
```

### 4. Reset cart (Delete all item(s) from the cart)

To reset all item(s) from the cart, an HTTP **DELETE** request should be made at `/api/cart` this endpoint.

This api call returns a `json` data as response in the following format:

```json
{
    "status": "success",
    "message": "All item(s) successfully deleted from cart. Total item(s) deleted: 2"
}
```

If the cart is empty then the response is like:

```json
{
    "status": "error",
    "message": "No product found. Cart is already empty."
}
```

### 5. Place an order

To make and order, an HTTP **POST** request should be made at `/api/order` this endpoint. The item(s) for the order is retrieved from the cart and `total price` is also calculated.

This api call returns a `json` data as response in the following format:

```json
{
    "status": "success",
    "id": "6101de8414805648138fce47",
    "message": "Order successfully placed. Order id: 6101de8414805648138fce47"
}
```

If the cart is empty then the response is like:

```json
{
    "status": "error",
    "message": "No product found in the cart. Cart is empty."
}
```