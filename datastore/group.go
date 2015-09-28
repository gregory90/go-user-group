package datastore

import (
	"database/sql"

	"bitbucket.org/pqstudio/go-user-group/model"

	. "bitbucket.org/pqstudio/go-user-group/db"
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
            userUID=unhex(?)
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

func GetAll(limit int, offset int) ([]model.Group, error) {
	rows, err := DB.Query(selectQuery+"ORDER BY createdAt DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}

	rs, err := getAll(rows)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func GetByUserUID(uid string, limit int, offset int) ([]model.Group, error) {
	rows, err := DB.Query(selectQuery+"WHERE userUID = unhex(?) ORDER BY createdAt DESC LIMIT ? OFFSET ?", uid, limit, offset)
	if err != nil {
		return nil, err
	}

	rs, err := getAll(rows)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func Get(uid string) (*model.Group, error) {
	r := &model.Group{}
	row := DB.QueryRow(selectQuery+"WHERE uid = unhex(?)", uid)

	err := scanSelectSingle(r, row)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetOneByUserUIDAndGroup(uid string, group string) (*model.Group, error) {
	r := &model.Group{}
	row := DB.QueryRow(selectQuery+"WHERE userUID = unhex(?) AND name = ?", uid, group)

	err := scanSelectSingle(r, row)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func Create(m *model.Group) error {
	stmt, err := DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = execInsert(m, stmt)

	return err
}

func Update(m *model.Group) error {
	stmt, err := DB.Prepare(updateQuery + "WHERE uid=unhex(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = execUpdate(m, stmt)
	return err
}

func Delete(uid string) error {
	stmt, err := DB.Prepare(deleteQuery + "WHERE uid=unhex(?)")
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
