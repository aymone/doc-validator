package main

import (
	"errors"
	"time"
)

type document struct {
	ID          string    `bson:"_id,omitempty" json:"documentNumber"`
	Variety     string    `bson:"variety" json:"variety"`
	Blacklisted bool      `bson:"blacklisted" json:"blacklisted"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
}

func (d *document) create() error {
	return getClient().C("documents").Insert(&d)
}

func (d *document) blacklist(ID string, status string) error {
	err := getClient().C("documents").FindId(ID).One(&d)
	if err != nil {
		return err
	}

	switch status {
	case "add":
		d.Blacklisted = true

	case "remove":
		d.Blacklisted = false

	default:
		return errors.New("Invalid Status, valid:add/remove")
	}

	return getClient().C("documents").UpdateId(d.ID, &d)
}

func (d *document) read() error {
	return getClient().C("documents").FindId(d.ID).One(&d)
}
