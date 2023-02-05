package models

import (
	"time"
	"example/myco-api/ent"
)

type SporeSyringeModel struct {
	ID int `json:"id"`
	Species string `json:"species"`
	Received time.Time `json:"received"`
	Supplier string `json:"supplier"`
}

func ToSporeSyringeModel(sporeSyringe *ent.SporeSyringe) SporeSyringeModel {
	sporeSyringeModel := SporeSyringeModel {
		ID: sporeSyringe.ID,
		Species: sporeSyringe.Species,
		Received: sporeSyringe.RecievedDate,
		Supplier: sporeSyringe.Supplier,
	}
	return sporeSyringeModel
}