package repositories

import (
	"fmt"

	"github.com/yassinebenaid/gomyadmin/database"
	"github.com/yassinebenaid/gomyadmin/database/models"
)

type DatabaseRepotitory struct {
	Connection database.Connection
}

func (r DatabaseRepotitory) List() ([]models.Database, error) {
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

	var databases []models.Database

	for result.Next() {
		var db models.Database
		if err := result.Scan(&db.Name, &db.DefaultCharsetName, &db.DefaultCollationName, &db.DefaultEncryption); err != nil {
			return nil, fmt.Errorf("database read error : %v", err)
		}
		databases = append(databases, db)
	}

	return databases, nil
}
