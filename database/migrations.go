package database

type Migrations struct {
	Table         string
	ColumnName    string
	MigrationName string // the name of the migration should be unique
	Query         string
}

func (m *Migrations) Check() (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM schema_migrations WHERE name = $1
		)
	`
	err := DB.QueryRow(query, m.MigrationName).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, err
}

func (m *Migrations) AddMigrationToDB() error {
	query := `
		INSERT INTO schema_migrations(name) VALUES ($1)
	`
	_, err := DB.Query(query, m.MigrationName)

	return err
}

func (m *Migrations) Apply() error {
	tx, err := DB.Begin()

	if err != nil {
		return err
	}

	// actual migration query
	if _, err := tx.Exec(m.Query); err != nil {
		tx.Rollback()
		return err
	}

	// adding the name of the migration to the schema_migrations DB
	if err := m.AddMigrationToDB(); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
