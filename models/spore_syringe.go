package models

import (
	"time"
)

type SporeSyringe struct {
	ID string `json:"id"`
	Species string `json:"species"`
	Received time.Time `json:"received"`
	Supplier string `json:"supplier"`
}