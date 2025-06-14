package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/tirlochanarora16/url_shortner/database"
)

type Urls struct {
	ID          string    `json:"id"`
	ShortCode   string    `json:"short_code"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewShortUrlBody struct {
	OriginalUrl string `json:"original_url" binding:"required"`
}

func (u *Urls) Save() (*Urls, error) {
	query := `
		INSERT INTO urls(short_code, original_url) VALUES ($1, $2)
		RETURNING id, short_code, original_url, created_at
	`
	rows, err := database.DB.Query(query, u.ShortCode, u.OriginalUrl)

	if err != nil {
		return &Urls{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&u.ID, &u.ShortCode, &u.OriginalUrl, &u.CreatedAt)
	}

	if err != nil {
		return &Urls{}, err
	}

	// return the newly created row
	return u, nil
}

func CheckUrlExists(originalUrl string) (*Urls, error) {
	query := "SELECT *  FROM urls WHERE original_url = $1"

	row := database.DB.QueryRow(query, originalUrl)

	var selectedRow Urls
	err := row.Scan(&selectedRow.ID, &selectedRow.ShortCode, &selectedRow.OriginalUrl, &selectedRow.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &Urls{}, err
	}

	return &selectedRow, nil
}

func CheckShortCode(shortCode string) (*Urls, error) {
	query := "SELECT * FROM urls WHERE short_code = $1"

	row := database.DB.QueryRow(query, shortCode)

	var selectedRow Urls
	err := row.Scan(&selectedRow.ID, &selectedRow.ShortCode, &selectedRow.OriginalUrl, &selectedRow.CreatedAt, &selectedRow.UpdatedAt)

	if err != nil {
		return &Urls{}, err
	}

	return &selectedRow, nil
}

func (u *Urls) Update() (*Urls, error) {
	query := `
		UPDATE urls SET original_url = $1, updated_at = NOW()
		WHERE short_code = $2
		RETURNING id, short_code, original_url
	`

	row, err := database.DB.Query(query, u.OriginalUrl, u.ShortCode)

	if err != nil {
		return &Urls{}, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&u.ID, &u.ShortCode, &u.OriginalUrl)
	}

	return u, nil
}

func (u *Urls) Delete() error {
	query := "DELETE FROM urls WHERE short_code = $1"
	result, err := database.DB.Exec(query, u.ShortCode)

	if err != nil {
		log.Println("Error deleting the url")
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println("could not delete the given row")
		return err
	}

	if rowsAffected == 0 {
		log.Println("cannot find a row with the given short code", u.ShortCode)
		return errors.New("cannot find a row with the given short code")
	}

	return nil
}
