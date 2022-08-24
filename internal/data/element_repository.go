package data

import (
	"context"
	"time"

	"github.com/cafaray/pkg/element"
)

type ElementRepository struct {
	Data *Data
}

func (er *ElementRepository) GetAll(ctx context.Context) ([]element.Element, error) {
	q := `
		SELECT id, user_id, key, value, created_at, updated_at FROM element;
	`
	rows, err := er.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var elems []element.Element
	for rows.Next() {
		var e element.Element
		rows.Scan(&e.ID, &e.UserID, &e.Key, &e.Value, &e.CreatedAt, &e.UpdatedAt)
		elems = append(elems, e)
	}
	return elems, nil
}
func (er *ElementRepository) GetOne(ctx context.Context, id uint) (element.Element, error) {
	q := `
	SELECT id, user_id, key, value, created_at, updated_at 
		FROM element
		WHERE id = $1;
	`
	row := er.Data.DB.QueryRowContext(ctx, q, id)
	var e element.Element
	err := row.Scan(&e.ID, &e.UserID, &e.Key, &e.Value, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		return element.Element{}, err
	}
	return e, nil
}
func (er *ElementRepository) GetByUser(ctx context.Context, userID uint) (element.Element, error) {
	q := `
	SELECT id, user_id, key, value, created_at, updated_at
		FROM element
		WHERE user_id = $1;
	`
	row := er.Data.DB.QueryRowContext(ctx, q, userID)
	var e element.Element
	err := row.Scan(&e.ID, &e.UserID, &e.Key, &e.Value, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		return element.Element{}, err
	}
	return e, nil
}
func (er *ElementRepository) Create(ctx context.Context, element *element.Element) error {
	q := `
		INSERT INTO elements (user_id, key, value, created_at, updated_at)
			VALUES ($1,$2,$3,$4,$5)
			RETURNING id;
	`
	stmt, err := er.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, element.UserID, element.Key, element.Value, time.Now(), time.Now())
	err = row.Scan(&element.ID)
	if err != nil {
		return err
	}
	return nil
}
func (er *ElementRepository) Update(ctx context.Context, id uint, e element.Element) error {
	q := `
	UPDATE element set user_id=$1, key=$2, value=$3, updated_at=$4 
		WHERE id=$5; 		
	`
	stmt, err := er.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, e.UserID, e.Key, e.Value, e.UpdatedAt, id)
	if err != nil {
		return err
	}
	return nil
}
func (er *ElementRepository) Delete(ctx context.Context, id uint) error {
	q := `
	DELETE FROM element WHERE id = $1;
	`
	stmt, err := er.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
