package service

import (
	"time"

	"bitbucket.org/pqstudio/go-user-group/datastore"
	"bitbucket.org/pqstudio/go-user-group/model"

	"bitbucket.org/pqstudio/go-webutils"
)

func Get(uid string) (*model.Group, error) {
	r, err := datastore.Get(uid)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetAll(limit int, offset int) ([]model.Group, error) {
	rs, err := datastore.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func GetByUserUID(uid string, limit int, offset int) ([]model.Group, error) {
	rs, err := datastore.GetByUserUID(uid, limit, offset)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func GetOneByUserUIDAndGroup(uid string, group string) (*model.Group, error) {
	r, err := datastore.GetOneByUserUIDAndGroup(uid, group)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func Create(m *model.Group) error {
	m.UID = utils.NewUUID()
	m.CreatedAt = time.Now().UTC()

	err := datastore.Create(m)
	if err != nil {
		return err
	}

	return nil
}

func Update(m *model.Group) error {
	err := datastore.Update(m)
	if err != nil {
		return err
	}

	return nil
}

func Delete(uid string) error {
	err := datastore.Delete(uid)
	if err != nil {
		return err
	}

	return nil
}
