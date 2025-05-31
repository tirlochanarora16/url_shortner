package database

var migrations = []map[string]string{
	{
		"table":      "urls",
		"columnName": "updated_at",
		"query":      AlterUrlTableUpdateAt,
	},
}
