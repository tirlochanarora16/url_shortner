package database

var MigrationsList = []Migrations{
	{
		Table:         "urls",
		ColumnName:    "updated_at",
		Query:         AlterUrlTableUpdateAt,
		MigrationName: "add_updated_at_1_06_25",
	},
	{
		Table:         "urls",
		ColumnName:    "access_count",
		Query:         AlterUrlTableAccessCount,
		MigrationName: "add_accessCount_15_06_25",
	},
	{
		Table:         "urls",
		ColumnName:    "short_code",
		Query:         ShortCodeUniqueConstraint,
		MigrationName: "short_code_constraint_06_07_25",
	},
}
