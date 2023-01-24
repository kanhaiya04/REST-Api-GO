package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func additem(q http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)
	q.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(q).Encode(newProfile)
}

func getitem(q http.ResponseWriter, r *http.Request) {
	json.NewEncoder(q).Encode(profiles)
}

func getitembyid(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(404)
		q.Write([]byte("Unable to conv into integer"))
		return
	}

	if id > len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("Id not found"))
		return
	}
	profile:= profiles[id]
	json.NewEncoder(q).Encode(profile)
}

func updateItem(q http.ResponseWriter, r *http.Request)  {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(404)
		q.Write([]byte("Unable to conv into integer"))
		return
	}

	if id > len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("Id not found"))
		return
	}
	var updatedProfile Profile
	json.NewDecoder(r.Body).Decode(&updatedProfile)
	q.Header().Set("Content-Type", "application/json")
	profiles[id]=updatedProfile
	json.NewEncoder(q).Encode(updatedProfile)
}

func deleteItem(q http.ResponseWriter, r *http.Request)  {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(404)
		q.Write([]byte("Unable to conv into integer"))
		return
	}

	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("Id not found"))
		return
	}
	profiles=append(profiles[:id], profiles[:id+1]...)
	q.WriteHeader(200)
	q.Write([]byte("Profile deleted"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profile", additem).Methods("POST")
	router.HandleFunc("/profile", getitem).Methods("GET")
	router.HandleFunc("/profile/{id}", getitembyid).Methods("GET")
	router.HandleFunc("/profile/{id}",updateItem).Methods("PUT")
	router.HandleFunc("/profile/{id}",deleteItem).Methods("DELETE")
	http.ListenAndServe(":5000", router)
}
