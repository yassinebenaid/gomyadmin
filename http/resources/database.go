package resources

import "github.com/yassinebenaid/gomyadmin/database/models"

type DatabaseResource struct {
	Data  []models.Database `json:"data"`
	Total int               `json:"total"`
}
