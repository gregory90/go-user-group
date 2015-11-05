package datastore

import (
	"database/sql"

	"github.com/gregory90/go-user-group/model"
)

const (
	table string = "groups"

	selectQuery string = `
        SELECT 
            lower(hex(uid)), 
            name, 
            lower(hex(userUID)), 
            createdAt 
        FROM ` + table + " "

	insertQuery string = `
        INSERT  ` + table + ` SET 
            uid=unhex(?),
            name=?,
            userUID=unhex(?),
            createdAt=?
             `

	updateQuery string = `
        UPDATE  ` + table + ` SET 
            name=?
             `
	deleteQuery string = `
        DELETE FROM  ` + table + ` `
)

func getAll(rows *sql.Rows) ([]model.Group, error) {
	rs := []model.Group{}
	defer rows.Close()
	r := &model.Group{}
	for rows.Next() {
		err := scanSelect(r, rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, *r)
	}
	err := rows.Err()

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func scanSelect(m *model.Group, rows *sql.Rows) error {
	err := rows.Scan(
		&m.UID,
		&m.Name,
		&m.UserUID,
		&m.CreatedAt,
	)
	return err
}

func scanSelectSingle(m *model.Group, row *sql.Row) error {
	err := row.Scan(
		&m.UID,
		&m.Name,
		&m.UserUID,
		&m.CreatedAt,
	)
	return err
}

func execInsert(m *model.Group, stmt *sql.Stmt) error {
	_, err := stmt.Exec(
		m.UID,
		m.Name,
		m.UserUID,
		m.CreatedAt,
	)

	return err
}

func execUpdate(m *model.Group, stmt *sql.Stmt) error {
	_, err := stmt.Exec(
		m.Name,
		m.UserUID,
		m.UID,
	)

	return err
}

func GetAll(tx *sql.Tx, limit int, offset int) ([]model.Group, error) {
	rows, err := tx.Query(selectQuery+"ORDER BY createdAt DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}

	rs, err := getAll(rows)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func GetByUserUID(tx *sql.Tx, uid string, limit int, offset int) ([]model.Group, error) {
	rows, err := tx.Query(selectQuery+"WHERE userUID = unhex(?) ORDER BY createdAt DESC LIMIT ? OFFSET ?", uid, limit, offset)
	if err != nil {
		return nil, err
	}

	rs, err := getAll(rows)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func Get(tx *sql.Tx, uid string) (*model.Group, error) {
	r := &model.Group{}
	row := tx.QueryRow(selectQuery+"WHERE uid = unhex(?)", uid)

	err := scanSelectSingle(r, row)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetOneByUserUIDAndGroup(tx *sql.Tx, uid string, group string) (*model.Group, error) {
	r := &model.Group{}
	row := tx.QueryRow(selectQuery+"WHERE userUID = unhex(?) AND name = ?", uid, group)

	err := scanSelectSingle(r, row)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func Create(tx *sql.Tx, m *model.Group) error {
	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = execInsert(m, stmt)

	return err
}

func Update(tx *sql.Tx, m *model.Group) error {
	stmt, err := tx.Prepare(updateQuery + "WHERE uid=unhex(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = execUpdate(m, stmt)
	return err
}

func Delete(tx *sql.Tx, uid string) error {
	stmt, err := tx.Prepare(deleteQuery + "WHERE uid=unhex(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uid)
	if err != nil {
		return err
	}

	return nil
}
