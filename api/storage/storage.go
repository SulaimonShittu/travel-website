package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"travel-website/api/models"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:trav2025@localhost:5432/travel?sslmode=disable"
		fmt.Println("[INFO] Using fallback hardcoded Postgres DSN")
	} else {
		fmt.Println("[INFO] Using DATABASE_URL from environment")
	}

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("cannot ping database: %v", err)
	}

	execMigrations()
}

func execMigrations() {
	contactTable := `
	CREATE TABLE IF NOT EXISTS contacts (
		id SERIAL PRIMARY KEY,
		full_name TEXT,
		email TEXT,
		phone_number TEXT,
		subject TEXT,
		comment TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	destTable := `
	CREATE TABLE IF NOT EXISTS destinations (
		id SERIAL PRIMARY KEY,
		name TEXT,
		ratings REAL,
		category TEXT,
		pricing DOUBLE PRECISION
	)`

	_, err := DB.Exec(contactTable)
	if err != nil {
		log.Fatalf("could not create contacts table: %v", err)
	}
	_, err = DB.Exec(destTable)
	if err != nil {
		log.Fatalf("could not create destinations table: %v", err)
	}
}

func SaveContact(c models.Contact) error {
	query := `INSERT INTO contacts (full_name, email, phone_number, subject, comment, created_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return DB.QueryRow(query, c.FullName, c.Email, c.Phone, c.Subject, c.Comment, c.CreatedAt).Scan(&c.ID)
}

func GetContact(id int64) (models.Contact, bool) {
	var c models.Contact
	query := `SELECT id, full_name, email, phone_number, subject, comment, created_at FROM contacts WHERE id = $1`
	err := DB.QueryRow(query, id).Scan(&c.ID, &c.FullName, &c.Email, &c.Phone, &c.Subject, &c.Comment, &c.CreatedAt)
	if err != nil {
		return c, false
	}
	return c, true
}

func GetAllContacts() []models.Contact {
	contacts := []models.Contact{}
	rows, err := DB.Query(`SELECT id, full_name, email, phone_number, subject, comment, created_at FROM contacts ORDER BY id`)
	if err != nil {
		log.Println("Error querying contacts:", err)
		return contacts
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Contact
		err := rows.Scan(&c.ID, &c.FullName, &c.Email, &c.Phone, &c.Subject, &c.Comment, &c.CreatedAt)
		if err == nil {
			contacts = append(contacts, c)
		}
	}
	return contacts
}

func SeedDestinations() {
	existing := GetAllDestinations()
	if len(existing) > 0 {
		return // already seeded
	}
	dests := []models.Destination{
		{Name: "Bali Beach Resort", Ratings: 4.5, Category: "Resort", Pricing: 299.99},
		{Name: "Swiss Alps Lodge", Ratings: 4.8, Category: "Lodge", Pricing: 399.99},
		{Name: "Tokyo Capsule Hotel", Ratings: 4.0, Category: "Budget", Pricing: 59.99},
	}
	for _, d := range dests {
		_, err := DB.Exec(`INSERT INTO destinations (name, ratings, category, pricing) VALUES ($1, $2, $3, $4)`,
			d.Name, d.Ratings, d.Category, d.Pricing)
		if err != nil {
			log.Println("Failed to insert destination:", err)
		}
	}
}

func GetAllDestinations() []models.Destination {
	dests := []models.Destination{}
	rows, err := DB.Query(`SELECT id, name, ratings, category, pricing FROM destinations ORDER BY id`)
	if err != nil {
		log.Println("Error querying destinations:", err)
		return dests
	}
	defer rows.Close()

	for rows.Next() {
		var d models.Destination
		err := rows.Scan(&d.ID, &d.Name, &d.Ratings, &d.Category, &d.Pricing)
		if err == nil {
			dests = append(dests, d)
		}
	}
	return dests
}
