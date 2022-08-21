package model

type Websites struct {
	List   []string          `json:"websites"`
	Status map[string]string `json:"-"`
}
