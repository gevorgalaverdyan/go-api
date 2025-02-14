package models

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type HealthFacility struct {
	Id           string `json:"Id"`
	IsOpen       bool   `json:"IsOpen"`
	Name         string `json:"Name"`
	PhoneNumber  string `json:"PhoneNumber"`
	StreetName   string `json:"StreetName"`
	StreetNumber string `json:"StreetNumber"`
	Type         string `json:"Type"`
}

// Custom UnmarshalJSON for HealthFacility
func (hf *HealthFacility) UnmarshalJSON(data []byte) error {
	type Alias HealthFacility
	aux := &struct {
		IsOpen string `json:"IsOpen"`
		*Alias
	}{
		Alias: (*Alias)(hf),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert string "true"/"false" to bool
	hf.IsOpen = strings.ToLower(aux.IsOpen) == "true"

	return nil
}

func GetAll() ([]HealthFacility, error) {
	var facilities []HealthFacility

	jsonFile, err := os.Open("/workspaces/go-api/db/HealthFacility.json")
	if err != nil {
		log.Panic("Opening failed", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &facilities)
	if err != nil {
		log.Panic("unmarshal failed", err)
		return nil, err
	}

	return facilities, nil
}

func GetById(id string) (*HealthFacility, error) {
	facilities, err := GetAll()

	if err != nil {
		log.Panic("getbyid", err)
		return nil, err
	}

	for _, f := range facilities {
		if f.Id == id {
			return &f, nil
		}
	}

	log.Print("No f found")
	return nil, err
}
