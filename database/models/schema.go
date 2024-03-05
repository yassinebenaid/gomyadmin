package models

type Schema struct {
	Name                    string `json:"name"`
	DefaultCharacterSetName string `json:"default_character_set_name"`
	DefaultCollationName    string `json:"default_collation_name"`
	DefaultEncryption       string `json:"default_encryption"`
}

type Collation struct {
	Name             string `json:"name"`
	CharacterSetName string `json:"character_set_name"`
	ID               string `json:"id"`
	IsDefault        string `json:"is_default"`
	IsCompiled       string `json:"is_compiled"`
	Sortlen          string `json:"sortlen"`
	PadAttribute     string `json:"pad_attribute"`
}
