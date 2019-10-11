package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	xmlBytes, err := getXML("http://data.bmkg.go.id/autogempa.xml")

	if err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		var Earthquake Infogempa
		xml.Unmarshal(xmlBytes, &Earthquake)

		// parsed earthquake
		fmt.Println(Earthquake)
	}
}
