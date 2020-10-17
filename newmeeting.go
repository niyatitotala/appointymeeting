package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMeeting(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	var tempmeet NewMeeting 
	_ = json.NewDecoder(request.Body).Decode(&tempmeet)
	tempmeet.def()
	if meet.Starttime != meet.Creationtime {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Invalid time. Please check again." }`))
		return
	}
	collection := client.Database("companydatabase").Collection("databasemeetings")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, tempmeet)
	tempmeet.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(response).Encode(tempmeet)
	fmt.Println(tempmeet)
}