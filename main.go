package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client 
var lock sync.Mutex

type Person struct {
	Name  string  `json:"firstname" bson:"firstname"`
	Email string  `json:"email" bson:"email"`
	RSVP  string  `json:"rsvp" bson:"rsvp"` 

type Meeting struct{
	ID	primitive.ObjectID `json:"_id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Participants []participants `json:"participants" bson:"participants"`
	Starttime    string `json:"starttime" bson:"starttime"`
	Endtime      string `json:"endtime" bson:"endtime"`
	Creationtime string `json:"creationtime" bson:"creationtime"`
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/meetings", MeetingHandler)
	http.HandleFunc("/articles/", GetParticipants)
	http.HandleFunc("/meeting/", GetMeetingwithID)
	http.ListenAndServe(":12345", nil)
}

