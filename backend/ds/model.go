package ds

import "time"

// ClientContact defines model for ClientContact.
type ClientContact struct {
	// contact name
	ClientName string `db:"client_name,size:128"`
	// email
	Email string `db:"email,size:128"`
	// message
	Message string `db:"message,size:1000"`
	// createdAt timestamp
	CreatedAt time.Time `db:"created_at"`
}
