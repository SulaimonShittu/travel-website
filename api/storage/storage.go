package storage

import (
	"sync"
	"travel-website/api/models"
)

var (
	contactStore           = make(map[int64]models.Contact)
	destinationStore       = make(map[int64]models.Destination)
	contactIDCounter int64 = 1
	destIDCounter    int64 = 1
	mutex            sync.Mutex
)

func SaveContact(c models.Contact) {
	contactStore[c.ID] = c
}

func GetContact(id int64) (models.Contact, bool) {
	c, ok := contactStore[id]
	return c, ok
}

func GetAllContacts() []models.Contact {
	contacts := make([]models.Contact, 0, len(contactStore))
	for _, c := range contactStore {
		contacts = append(contacts, c)
	}
	return contacts
}

func NextContactID() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	id := contactIDCounter
	contactIDCounter++
	return id
}

func SeedDestinations() {
	dests := []models.Destination{
		{ID: 1, Name: "Bali Beach Resort", Ratings: 4.5, Category: "Resort", Pricing: 299.99},
		{ID: 2, Name: "Swiss Alps Lodge", Ratings: 4.8, Category: "Lodge", Pricing: 399.99},
		{ID: 3, Name: "Tokyo Capsule Hotel", Ratings: 4.0, Category: "Budget", Pricing: 59.99},
	}

	for _, d := range dests {
		destinationStore[d.ID] = d
		if d.ID >= destIDCounter {
			destIDCounter = d.ID + 1
		}
	}
}

func GetAllDestinations() []models.Destination {
	dests := make([]models.Destination, 0, len(destinationStore))
	for _, d := range destinationStore {
		dests = append(dests, d)
	}
	return dests
}
