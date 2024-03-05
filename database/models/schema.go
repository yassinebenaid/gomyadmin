package models

type Schema struct {
	Name                 string `json:"name"`
	DefaultCharsetName   string `json:"default_charset"`
	DefaultCollationName string `json:"default_collation_name"`
	DefaultEncryption    string `json:"default_encryption"`
}
