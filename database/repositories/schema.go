package repositories

import (
	"database/sql"
	"fmt"

	"github.com/yassinebenaid/gomyadmin/database/models"
)

type SchemaRepotitory struct {
	Connection *sql.DB
}

func (r SchemaRepotitory) ListSchemas() ([]models.Schema, error) {
	result, err := r.Connection.Query(`
		SELECT SCHEMA_NAME, DEFAULT_CHARACTER_SET_NAME, DEFAULT_COLLATION_NAME,DEFAULT_ENCRYPTION 
		FROM information_schema.SCHEMATA
	`)
	if err != nil {
		return nil, fmt.Errorf("database query error : %v", err)
	}

	var schemas []models.Schema

	for result.Next() {
		var s models.Schema
		if err := result.Scan(&s.Name, &s.DefaultCharacterSetName, &s.DefaultCollationName, &s.DefaultEncryption); err != nil {
			return nil, fmt.Errorf("database read error : %v", err)
		}
		schemas = append(schemas, s)
	}

	return schemas, nil
}

func (r SchemaRepotitory) ListCollations() ([]models.Collation, error) {
	result, err := r.Connection.Query(`
		SELECT COLLATION_NAME, CHARACTER_SET_NAME, ID, IS_DEFAULT, IS_COMPILED, SORTLEN, PAD_ATTRIBUTE 
		FROM information_schema.COLLATIONS
	`)
	if err != nil {
		return nil, fmt.Errorf("database query error : %v", err)
	}

	var collations []models.Collation

	for result.Next() {
		var s models.Collation
		if err := result.Scan(&s.Name, &s.CharacterSetName, &s.ID, &s.IsDefault, &s.IsCompiled, &s.Sortlen, &s.PadAttribute); err != nil {
			return nil, fmt.Errorf("database read error : %v", err)
		}
		collations = append(collations, s)
	}

	return collations, nil
}
