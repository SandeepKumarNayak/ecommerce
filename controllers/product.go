package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sandeepkumarnayak/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""
const dbName = "ecommerce"
const colName = "product"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB Connection is successful")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection is ready")
}

func addProduct(product models.Product) {
	myProduct, err := collection.InsertOne(context.Background(), product)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted Successfully", myProduct.InsertedID)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var product models.Product

	_ = json.NewDecoder(r.Body).Decode(&product)
	addProduct(product)
	json.NewEncoder(w).Encode(product)
}

//for getting all products

func getAllProducts() []primitive.M {
	// cursor return
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var products []primitive.M
	for cur.Next(context.Background()) {
		var product primitive.M
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	defer cur.Close(context.Background())
	return products
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatin/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	products := getAllProducts()
	json.NewEncoder(w).Encode(products)
}

// get Product by ID
func getProductById(product_id string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(product_id)

	filter := bson.M{"_id": id}
	cur := collection.FindOne(context.Background(), filter)

	var product primitive.M
	err := cur.Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	return product
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	product := getProductById(params["id"])

	json.NewEncoder(w).Encode(product)
}

//update product

// delete a product
func deleteProductById(product_id string) {
	id, _ := primitive.ObjectIDFromHex(product_id)

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total Deleted Item", deleteCount.DeletedCount)
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteProductById(params["id"])
	json.NewEncoder(w).Encode("Product Deleted Successfully.")
}
