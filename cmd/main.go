package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	contact "contact.com"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

var mh *contact.MongoHandler

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		//Journey CMS
		r.Get("/journey", GetAllJourney)
		r.Get("/journey/{code}", GetDetailJourney)
	})
	// r.Route("/journey", func(r chi.Router) {
	// 	r.Get("/", getAllJourney) //GET /journey
	// 	r.Get("/{code}", getJourney) //GET /journey/0147344454
	// 	// r.Post("/", addContact)                   //POST /journey
	// 	// r.Put("/{phonenumber}", updateContact)    //PUT /journey/0147344454
	// 	// r.Delete("/{phonenumber}", deleteContact) //DELETE /journey/0147344454
	// })
	return r
}

func main() {
	mongoDbConnection := "mongodb://localhost:27017"
	mh = contact.NewHandler(mongoDbConnection)
	r := registerRoutes()
	log.Printf("Running on Debug Mode: On at host [127.0.0.1:3060]")
	log.Fatal(http.ListenAndServe(":3060", r))
}

//GetAllJourney ...
func GetAllJourney(w http.ResponseWriter, r *http.Request) {
	journey := mh.Get(bson.M{})
	json.NewEncoder(w).Encode(journey)
}

// GetDetailJourney ...
func GetDetailJourney(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	journey := &contact.JourneyPlan{}
	err := mh.GetOne(journey, bson.M{"code": code})
	if err != nil {
		http.Error(w, fmt.Sprintf("Journey with code: %s not found", code), 404)
		return
	}
	json.NewEncoder(w).Encode(journey)
}

// func addContact(w http.ResponseWriter, r *http.Request) {
// 	existingContact := &contact.JourneyPlan{}
// 	var contact contact.JourneyPlan
// 	json.NewDecoder(r.Body).Decode(&contact)
// 	contact.CreatedOn = time.Now()
// 	err := mh.GetOne(existingContact, bson.M{"phoneNumber": contact.PhoneNumber})
// 	if err == nil {
// 		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s already exist", contact.PhoneNumber), 400)
// 		return
// 	}
// 	_, err = mh.AddOne(&contact)
// 	if err != nil {
// 		http.Error(w, fmt.Sprint(err), 400)
// 		return
// 	}
// 	w.Write([]byte("Contact created successfully"))
// 	w.WriteHeader(201)
// }

// func deleteContact(w http.ResponseWriter, r *http.Request) {
// 	existingContact := &contact.Contact{}
// 	phoneNumber := chi.URLParam(r, "phonenumber")
// 	if phoneNumber == "" {
// 		http.Error(w, http.StatusText(404), 404)
// 		return
// 	}
// 	err := mh.GetOne(existingContact, bson.M{"phoneNumber": phoneNumber})
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s does not exist", phoneNumber), 400)
// 		return
// 	}
// 	_, err = mh.RemoveOne(bson.M{"phoneNumber": phoneNumber})
// 	if err != nil {
// 		http.Error(w, fmt.Sprint(err), 400)
// 		return
// 	}
// 	w.Write([]byte("Contact deleted"))
// 	w.WriteHeader(200)
// }

// func updateContact(w http.ResponseWriter, r *http.Request) {
// 	phoneNumber := chi.URLParam(r, "phonenumber")
// 	if phoneNumber == "" {
// 		http.Error(w, http.StatusText(404), 404)
// 		return
// 	}
// 	contact := &contact.Contact{}
// 	json.NewDecoder(r.Body).Decode(contact)
// 	_, err := mh.Update(bson.M{"phoneNumber": phoneNumber}, contact)
// 	if err != nil {
// 		http.Error(w, fmt.Sprint(err), 400)
// 		return
// 	}
// 	w.Write([]byte("Contact update successful"))
// 	w.WriteHeader(200)
// }
