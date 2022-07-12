package app



import (

    "context"

    "fmt"

    "log"



    "go.mongodb.org/mongo-driver/mongo"

    "go.mongodb.org/mongo-driver/mongo/options"

)



func db() *mongo.Client {

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")



    //Connect to mongoDB

   

    client, err := mongo.Connect(context.TODO(), clientOptions)



    if err != nil {

        log.Fatal(err)

    }



    fmt.Println("Connected to MongoDB")



    //Check the Connection

    err = client.Ping(context.TODO(), nil)



    if err != nil {

        log.Fatal(err)

    }



    fmt.Println("Connection is established")

    return client



}