package service

import (
	"context"
	"example/myco-api/ent"
	"example/myco-api/ent/sporesyringe"
	"example/myco-api/config"
)

type SporeSyringOps struct {
	ctx context.Context
	client *ent.Client
}

func NewSporeSyringOps(ctx context.Context)  *SporeSyringOps {
	return &SporeSyringOps{
		ctx: ctx,
		client: config.GetClient(),
	}
}

func (r *SporeSyringOps) SporeSyringGetByID(id int) (*ent.SporeSyringe, error) {

	sporeSyring, err := r.client.SporeSyringe.Query().
		Where(
		ent.SporeSyringe.ID(id),
	).Only(r.ctx)
	if err != nil {
		return nil, err
	}

	return sporeSyring, nil
}

func (r *SporeSyringOps) SporeSyringCreate(newSporeSyring ent.SporeSyringe) (*ent.SporeSyring, error) {

	user, err := r.client.User.Query().Where(
		user.ID(newSporeSyring.Edges.Owner.ID),
	).Only(r.ctx)
	if err != nil {
		return nil, err
	}

	sporeSyring, err := r.client.SporeSyring.Create().
		SetName(newSporeSyring.Name).
		SetOwner(user).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return sporeSyring, nil
}

func (r *SporeSyringOps) SporeSyringUpdate(sporeSyring ent.SporeSyring)(*ent.SporeSyring, error) {

	updatedSporeSyring, err := r.client.SporeSyring.UpdateOneID(sporeSyring.ID).
		SetIsDone(sporeSyring.IsDone).
		SetName(sporeSyring.Name).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedSporeSyring, nil
}

func (r *SporeSyringOps) SporeSyringDelete(id int) (int, error) {

	err := r.client.SporeSyring.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}