package database

const CreateUrlTable = `
		CREATE TABLE IF NOT EXISTS urls (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			short_code VARCHAR(10) UNIQUE NOT NULL,
			original_url TEXT NOT NULL,
			created_at TIMESTAMP  DEFAULT NOW(),
			updated_at TIMESTAMP  DEFAULT NOW()
		)
`

const CreateSchemaMigrationTable = `
		CREATE TABLE IF NOT EXISTS schema_migrations (
    	name TEXT PRIMARY KEY,
    	applied_at TIMESTAMP DEFAULT NOW()
	)
`

const AlterUrlTableUpdateAt = "ALTER TABLE urls ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT NOW()"

const AlterUrlTableAccessCount = "ALTER TABLE urls ADD COLUMN IF NOT EXISTS access_count INTEGER DEFAULT 0"
