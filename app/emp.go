package app

import (
    "context"
    "encoding/json"
    "fmt"
    //"log"
    "net/http"
    "reflect"
    "time"

    //"github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson"
  //  "go.mongodb.org/mongo-driver/bson/primitive"
    //"golang.org/x/crypto/bcrypt"
)

type emp struct {
	emp_id    int64  `json:"emp_id"   bson:"emp_id"`
    FirstName string `json:"firstname" bson:"firstname"`
    LastName  string `json:"lastname" bson:"lastname"`
    DOB       string `json:"dob"       bson:"DOB"`
    Gender    string `json:"Gender"    bson:"Gender"`
    //State     string `json:"State"    bson:"State"`
    //City      string `json:"City"    bson:"City"`
    PhoneNO   int64  `json:"phoneNO"   bson:"PhoneNO"`
    Email     string `json:"email" bson:"email"`
    
}
type positiveResponse struct {
    StatusCode    int    `json:"StatusCode" bson:"StatusCode"`
    Status        bool   `json:"Status" bson:"Status"`
    CustomMessage string `json:"CustomMessage" bson:"CustomMessage"`
}

type negativeResponse struct {
    ErrorMessage  string `json:"ErrorMessage" bson:"ErrorMessage"`
    StatusCode    int    `json:"StatusCode" bson:"StatusCode"`
    Status        bool   `json:"Status" bson:"Status"`
    CustomMessage string `json:"CustomMessage" bson:"CustomMessage"`
}

// code complient from ganga


func CreateProfile(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var person emp
    json.NewDecoder(r.Body).Decode(&person)
    //person.Password = getHash([]byte(person.Password))
    collection := db().Database("employee").Collection("details")
    fmt.Println("Collection type:", reflect.TypeOf(collection), "\n")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    err := collection.FindOne(ctx, bson.M{"email": person.Email}).Decode(&person)
    if err != nil {
        result, _ := collection.InsertOne(ctx, person)

        fmt.Println("Document Inserted:", result)
        msg := positiveResponse{
            StatusCode:    200,
            Status:        true,
            CustomMessage: "Document Inserted Successfully",
        }
        json.NewEncoder(w).Encode(msg)

    } else {
        msg2 := negativeResponse{
            ErrorMessage:  "nil",
            StatusCode:    200,
            Status:        false,
            CustomMessage: "Mail Id exist Already",
        }

        json.NewEncoder(w).Encode(msg2)
	}
}

