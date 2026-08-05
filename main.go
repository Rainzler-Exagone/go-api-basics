package main

import (
	"fmt"
	"net/http"
	"encoding/json"

)


func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}


var users = []user{
	{Id: 1, Name: "John Doe"},
	{Id: 2, Name: "Jane Smith"},
}

func getUsers(w http.ResponseWriter, r  *http.Request){

	w.Header().Set("Content-type","application/json")

	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request){

	var newUser user

	err:=json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w,"Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Received new user:", newUser)

	newUser.Id = len(users) + 1

	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}




func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

		case http.MethodGet:
			getUsers(w, r)
	
	    case http.MethodPost:
			createUser(w, r)
	
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main(){
	
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/users", UsersHandler)
	fmt.Println("Server running on http://localhost:8080")	
	http.ListenAndServe(":8080", nil)



}