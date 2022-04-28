package cases

import (
	"assignment3/models"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var startTime time.Time

type UseCase struct{}

func NewCase() models.StatusCase {
	return &UseCase{}
}

func (u *UseCase) UpdateStatus() (models.Danger, error) {
	var data models.Data
	var statusData models.Danger
	jsonFile, err := os.Open("value.json")
	if err != nil {
		fmt.Println(err)
	}

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json, error: %v", err)
		return statusData, err
	}

	_ = json.Unmarshal(jsonData, &data)

	if maketime() > 15*time.Second {
		data.Data.Water = rand.Intn(100)
		data.Data.Wind = rand.Intn(100)
		jsonString, _ := json.Marshal(data)
		_ = os.WriteFile("value.json", jsonString, os.ModePerm)
		startTime = time.Now()
	}

	statusData.Water = data.Data.Water
	statusData.Wind = data.Data.Wind

	if data.Data.Water <= 5 {
		statusData.StatusWater = "Aman"
	} else if data.Data.Water > 5 && data.Data.Water <= 8 {
		statusData.StatusWater = "Siaga"
	} else {
		statusData.StatusWater = "Bahaya"
	}

	if data.Data.Wind <= 6 {
		statusData.StatusWind = "Aman"
	} else if data.Data.Wind > 6 && data.Data.Wind <= 15 {
		statusData.StatusWind = "Siaga"
	} else {
		statusData.StatusWind = "Bahaya"
	}

	fmt.Println(statusData)

	return statusData, nil

}

func maketime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}
