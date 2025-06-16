package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/tirlochanarora16/url_shortner/database"
)

type Urls struct {
	ID          string    `json:"id"`
	ShortCode   string    `json:"short_code"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessCount int       `json:"access_count"`
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

func IsValidUrl(urlString string) bool {
	u, err := url.ParseRequestURI(urlString)

	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	return true
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
	err := row.Scan(&selectedRow.ID, &selectedRow.ShortCode, &selectedRow.OriginalUrl, &selectedRow.CreatedAt, &selectedRow.UpdatedAt, &selectedRow.AccessCount)

	if err != nil {
		return &Urls{}, err
	}

	return &selectedRow, nil
}

func (u *Urls) UpdateUrl(fields map[string]any) (*Urls, error) {
	if len(fields) == 0 {
		log.Println("No fields provided to update")
		return &Urls{}, errors.New("No fields provided to update")
	}

	var i int = 1
	columnInput := ""
	var values = []any{}

	skipUpdate := false

	for key, value := range fields {
		key = strings.TrimSpace(key)
		if key == "" {
			if len(fields) == 1 {
				skipUpdate = true
				break
			}
			continue
		}
		switch v := value.(type) {
		case string:
			value = strings.TrimSpace(v)
			if value == "" {
				if len(fields) == 1 {
					skipUpdate = true
					break
				}
				continue
			}
		}

		columnInput += fmt.Sprintf("%s = %s, ", key, fmt.Sprintf("$%d", i))
		values = append(values, value)
		i++
	}

	if skipUpdate {
		return &Urls{}, errors.New("column name and value cannot be empty")
	}

	values = append(values, u.ID) // thereby last "i" value would be the ID of the url for matching

	columnInput = strings.TrimSpace(columnInput)
	columnInput = strings.TrimSuffix(strings.TrimSpace(columnInput), ",")

	query := fmt.Sprintf("UPDATE urls SET %s, updated_at = NOW() WHERE id = $%d RETURNING *", columnInput, i)

	err := database.DB.QueryRow(query, values...).Scan(
		&u.ID,
		&u.ShortCode,
		&u.OriginalUrl,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.AccessCount,
	)

	if err != nil {
		log.Println("Updating urls table failed")
		return &Urls{}, err
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
