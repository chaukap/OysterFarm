package models

import (
	"time"
)

type GrainJar struct {
	ID string `json:"id"`
	Grain string `json:"grain"`
	InnoculationDate time.Time `json:"innoculation_date"`
	HarvestDate time.Time `json:"harvest_date"`
	SporeSyringe SporeSyringe `json:"spore_syringe"`
}