package resources

import "github.com/yassinebenaid/gomyadmin/database/models"

type SchemaResource struct {
	Data  []models.Schema `json:"data"`
	Total int             `json:"total"`
}
