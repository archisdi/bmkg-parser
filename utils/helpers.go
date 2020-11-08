package utils

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// GetXMLFromURL ...
func GetXMLFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

// CalculateEuclideanDistance ...
func CalculateEuclideanDistance(pointA string, pointB string) float64 {
	x := strings.Split(pointA, ",")
	y := strings.Split(pointB, ",")

	xA, _ := strconv.ParseFloat(x[0], 10)
	xB, _ := strconv.ParseFloat(x[1], 10)

	yA, _ := strconv.ParseFloat(y[0], 10)
	yB, _ := strconv.ParseFloat(y[1], 10)

	return math.Sqrt(math.Pow((xA-yA), 2) + math.Pow((xB-yB), 2))
}

// SpaceFieldsJoin ...
func SpaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}
