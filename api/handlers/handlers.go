package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
	"travel-website/api/models"
	"travel-website/api/storage"

	"github.com/go-chi/chi/v5"
)

var mutex sync.Mutex

func ContactRoutes(r chi.Router) {
	r.Post("/", createContactHandler)
	r.Get("/{id}", getContactByIDHandler)
	r.Get("/", getAllContactsHandler)
}

func DestinationRoutes(r chi.Router) {
	r.Get("/", getAllDestinationsHandler)
}

func createContactHandler(w http.ResponseWriter, r *http.Request) {
	var c models.Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	c.CreatedAt = time.Now()

	if err := storage.SaveContact(c); err != nil {
		http.Error(w, "could not save contact", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func getContactByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	contact, ok := storage.GetContact(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func getAllContactsHandler(w http.ResponseWriter, r *http.Request) {
	contacts := storage.GetAllContacts()
	json.NewEncoder(w).Encode(contacts)
}

func getAllDestinationsHandler(w http.ResponseWriter, r *http.Request) {
	dests := storage.GetAllDestinations()
	json.NewEncoder(w).Encode(dests)
}
