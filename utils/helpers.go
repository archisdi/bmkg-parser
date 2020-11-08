package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"github.com/hashicorp/go-multierror"
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

// StringToCoordinate ...
func StringToCoordinate(coordinate string) (float64, float64, error) {
	x := strings.Split(SpaceFieldsJoin(coordinate), ",")

	if len(x) != 2 {
		return 0,0, errors.New("invalid coordinate format")
	}

	var errColl error
	xA, errA := strconv.ParseFloat(x[0], 10)
	if errA != nil {
		errColl = multierror.Append(errColl, errA)
	}

	xB, errB := strconv.ParseFloat(x[1], 10)
	if errB != nil {
		errColl = multierror.Append(errColl, errB)
	}

	return xA, xB, errColl
}

// CalculateEuclideanDistance ...
func CalculateEuclideanDistance(xA, xB, yA, yB float64) float64 {
	return math.Sqrt(math.Pow(xA-yA, 2) + math.Pow(xB-yB, 2))
}

// SpaceFieldsJoin ...
func SpaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}
