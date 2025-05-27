package models

import (
	"database/sql"
	"time"

	"github.com/tirlochanarora16/url_shortner/database"
)

type Urls struct {
	ID          string
	ShortCode   string
	OriginalUrl string
	DateTime    time.Time
}

func CheckUrlExists(originalUrl string) (*Urls, error) {
	query := "SELECT *  FROM urls WHERE original_url = $1"

	row := database.DB.QueryRow(query, originalUrl)

	var selectedRow Urls
	err := row.Scan(&selectedRow.ID, &selectedRow.ShortCode, &selectedRow.OriginalUrl, &selectedRow.DateTime)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &Urls{}, err
	}

	return &selectedRow, nil
}
