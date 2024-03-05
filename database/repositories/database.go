package repositories

import (
	"fmt"

	"github.com/yassinebenaid/gomyadmin/database"
)

type DatabaseRepotitory struct {
	Connection database.Connection
}

func (r DatabaseRepotitory) List() ([]string, error) {
	conn, err := database.Connect(r.Connection)
	if err != nil {
		return nil, fmt.Errorf("database connection error : %v", err)
	}

	result, err := conn.Query("show databases")
	if err != nil {
		return nil, fmt.Errorf("database query error : %v", err)
	}

	var databases []string

	for result.Next() {
		var db_name string
		if err := result.Scan(&db_name); err != nil {
			return nil, fmt.Errorf("database read error : %v", err)
		}
		databases = append(databases, db_name)
	}

	return databases, nil
}
