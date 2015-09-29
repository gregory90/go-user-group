package service

import (
	"database/sql"
	"time"

	"bitbucket.org/pqstudio/go-user-group/datastore"
	"bitbucket.org/pqstudio/go-user-group/model"

	"bitbucket.org/pqstudio/go-webutils"
)

func Get(tx *sql.Tx, uid string) (*model.Group, error) {
	r, err := datastore.Get(tx, uid)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetAll(tx *sql.Tx, limit int, offset int) ([]model.Group, error) {
	rs, err := datastore.GetAll(tx, limit, offset)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func GetByUserUID(tx *sql.Tx, uid string, limit int, offset int) ([]model.Group, error) {
	rs, err := datastore.GetByUserUID(tx, uid, limit, offset)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func GetOneByUserUIDAndGroup(tx *sql.Tx, uid string, group string) (*model.Group, error) {
	r, err := datastore.GetOneByUserUIDAndGroup(tx, uid, group)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func Create(tx *sql.Tx, m *model.Group) error {
	m.UID = utils.NewUUID()
	m.CreatedAt = time.Now().UTC()

	err := datastore.Create(tx, m)
	if err != nil {
		return err
	}

	return nil
}

func Update(tx *sql.Tx, m *model.Group) error {
	err := datastore.Update(tx, m)
	if err != nil {
		return err
	}

	return nil
}

func Delete(tx *sql.Tx, uid string) error {
	err := datastore.Delete(tx, uid)
	if err != nil {
		return err
	}

	return nil
}
