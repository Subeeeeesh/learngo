package main



import (

    "learngo/app"

    "fmt"

    "log"

    "net/http"



    "github.com/gorilla/mux"

)



func main() {

    route := mux.NewRouter()

    s := route.PathPrefix("/api").Subrouter() //Base path



    fmt.Println("Server set up done")



    //Routes

    s.HandleFunc("/CreateProfile", app.CreateProfile).Methods("POST")

    //s.HandleFunc("/loginProfile", app.UserLogin).Methods("POST")

    //s.HandleFunc("/updateProfile", app.UpdateProfile).Methods("POST")

    //s.HandleFunc("/category", app.Categories).Methods("POST")



    log.Fatal(http.ListenAndServe(":8080", s)) // Run Server

}