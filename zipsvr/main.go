package main

import (
	"log"
	"info344-in-class/zipsvr/models"
)

func main() {
	zips, err := models.LoadZips("zip.csv")
	if err != nil {
		log.Fatal("error loading zips: %v", err)
	}
	log.Printf("load %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := string.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

}
