package models

type (
	GetAllDictionary struct {
		DctSlug string `json:"dct_slug"`
		DctName string `json:"dct_name"`
		DctDesc string `json:"dct_desc"`
	}
)
