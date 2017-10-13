package main

import (
	"log"
	"github.com/aswfan/info344-in-class/zipsvr/models"
	"strings"
)

func main() {
	zips, err := models.LoadZip("zip.csv")
	if err != nil {
		log.Fatal("error loading zips: %v", err)
	}
	log.Printf("load %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

}
