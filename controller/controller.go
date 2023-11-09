package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	model "mogoIn/Models"
	"net/http"

	// "command-line-arguments/home/ankit/Desktop/Golang/GoMongo/Models/model.go"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



const dbName = "netflix"
const collectionName = "watchlist"


var collection  *mongo.Collection



func init()  {
	var clientOption = options.Client().ApplyURI("mongodb://localhost:27017")	

	client , err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB",client)

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance created")


}



func insertOneMovie(movie model.Show){

	inserted, err := collection.InsertOne(context.Background(), movie)
	
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted a single document: ", inserted.InsertedID)


}

func updateOne(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	updated, err := collection.UpdateOne(context.Background(), filter, update)

	if err!= nil{
		panic(err)
	}

	fmt.Println("Updated a single document: ",updated.UpsertedID)
	

}


func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId);

	filter := bson.M{"_id": id}

	deleted, err := collection.DeleteOne(context.Background(), filter,)
	if err !=nil{
		panic(err)
	}
	fmt.Println("Deleted a single document: ",deleted.DeletedCount)


}

func deleteMany(){
	collection.DeleteMany(context.Background(), bson.M{})

	// var movies []bson.M

	// for 

}

func findMany() []bson.M{
	cur,_ := collection.Find(context.Background(), bson.M{})

	var movies []bson.M

	for cur.Next(context.Background()){
		var movie bson.M

		err := cur.Decode(&movie)
		if err!= nil{
			panic(err)
		} 
		movies = append(movies, movie)


	}
	defer cur.Close(context.Background())
	return movies
}

func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	movies := findMany()
	
	json.NewEncoder(w).Encode(movies)


}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Show

	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)


}
func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOne(params["id"])

	json.NewEncoder(w).Encode("Movie Updated")
	

}