package requests

import (
	"fmt"
	"regexp"
)

type CreateSchemaRequest struct {
	Name      string `json:"name"`
	Collation string `json:"collation"`
}

func (r CreateSchemaRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf(`the "name" attribute is required`)
	}
	if !regexp.MustCompile(`^[A-Za-z0-9_]+$`).MatchString(r.Name) {
		return fmt.Errorf(`the "name" attribute format is invalid`)
	}
	if r.Collation == "" {
		return fmt.Errorf(`the "collation" attribute is required`)
	}

	return nil
}
