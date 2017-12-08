package main

import (
	"errors"
	"time"
)

type documentRequest struct {
	ID string `json:"id"`
}

// document represents mongodb collection model
type document struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Variety     string    `bson:"variety" json:"variety"`
	Blacklisted bool      `bson:"blacklisted" json:"blacklisted"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
}

// set status update document status and updatedAt field
func (d *document) setStatus(status string) error {
	switch status {
	case "add":
		d.Blacklisted = true

	case "remove":
		d.Blacklisted = false

	default:
		return errors.New("Invalid Status, valid:add/remove")
	}

	d.UpdatedAt = time.Now()

	return nil
}
