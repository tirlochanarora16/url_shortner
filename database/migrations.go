package database

var migrations = []map[string]string{
	{
		"table":      "urls",
		"columnName": "updated_at",
		"query":      "ALTER TABLE urls ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT NOW()",
	},
}
