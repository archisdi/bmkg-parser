package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func main() {
	xmlBytes, err := getXML("http://data.bmkg.go.id/gempaterkini.xml")

	if err != nil {
		log.Printf("Failed to get XML: %v", err)
		os.Exit(1)
	}

	var Earthquakes EarthquakeList
	xml.Unmarshal(xmlBytes, &Earthquakes)

	// parsed earthquake
	for _, earthquake := range Earthquakes.Gempa {
		fmt.Println(earthquake)
	}
}
