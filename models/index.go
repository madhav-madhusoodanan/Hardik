package models

type AddEntry struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type DeleteEntry struct {
	Key string `json:"key"`
}