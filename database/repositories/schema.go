package repositories

import (
	"fmt"

	"github.com/yassinebenaid/gomyadmin/database"
	"github.com/yassinebenaid/gomyadmin/database/models"
)

type SchemaRepotitory struct {
	Connection database.Connection
}

func (r SchemaRepotitory) List() ([]models.Schema, error) {
	conn, err := database.Connect(r.Connection)
	if err != nil {
		return nil, fmt.Errorf("database connection error : %v", err)
	}

	result, err := conn.Query(`
		SELECT SCHEMA_NAME, DEFAULT_CHARACTER_SET_NAME, DEFAULT_COLLATION_NAME,DEFAULT_ENCRYPTION 
		FROM information_schema.SCHEMATA
	`)
	if err != nil {
		return nil, fmt.Errorf("database query error : %v", err)
	}

	var schemas []models.Schema

	for result.Next() {
		var s models.Schema
		if err := result.Scan(&s.Name, &s.DefaultCharsetName, &s.DefaultCollationName, &s.DefaultEncryption); err != nil {
			return nil, fmt.Errorf("database read error : %v", err)
		}
		schemas = append(schemas, s)
	}

	return schemas, nil
}
