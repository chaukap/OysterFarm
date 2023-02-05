package models

import (
	"time"
	"example/myco-api/ent"
)

type GrainJarModel struct {
	ID int `json:"id"`
	Grain string `json:"grain"`
	InnoculationDate time.Time `json:"innoculation_date"`
	HarvestDate time.Time `json:"harvest_date"`
	SporeSyringeModel SporeSyringeModel `json:"spore_syringe"`
}

func ToGrainJarModel(grainJar *ent.GrainJar, sporeSyringe *ent.SporeSyringe) GrainJarModel {
	var syringeModel = ToSporeSyringeModel(sporeSyringe)
	model := GrainJarModel {
		ID: grainJar.ID,
		Grain: grainJar.Grain,
		InnoculationDate: grainJar.InnoculationDate,
		HarvestDate: grainJar.HarvestDate,
		SporeSyringeModel: syringeModel,
	}
	return model
}