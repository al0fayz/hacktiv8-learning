package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"sync"
	"time"
)

type Data struct {
	Wind  int64 `json:"wind"`
	Water int64 `json:"water"`
}
type Status struct {
	Status Data `json:"status"`
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go AutoReload(&wg) //use goroutine biar keren padahal pake loop saja bisa
	}

	wg.Wait()
}

func AutoReload(wg *sync.WaitGroup) {
	defer wg.Done() //close
	updateData()
	status, err := OpenAndReadFile()
	if err != nil {
		fmt.Println(err.Error())
	}
	logicWater(int(status.Status.Water))
	logicWind(int(status.Status.Wind))
	time.Sleep(15 * time.Second)
}
func logicWind(wind int) {
	//logic wind
	var windResult string
	var messages = fmt.Sprintf("Wind value: %d", wind)
	if wind < 6 {
		windResult = "Wind Status aman"
	} else if wind > 6 && wind <= 15 {
		windResult = "Wind Status siaga"
	} else if wind > 15 {
		windResult = "Wind Status bahaya"
	}
	fmt.Println(messages)
	fmt.Println(windResult)
}
func logicWater(water int) {
	//logic water
	var messages = fmt.Sprintf("Water value: %d", water)
	var waterResult string
	if water < 5 {
		waterResult = "Water Status aman"
	} else if water > 5 && water <= 8 {
		waterResult = "Water Status siaga"
	} else if water > 8 {
		waterResult = "Water Status Bahaya"
	}
	fmt.Println(messages)
	fmt.Println(waterResult)
}
func updateData() {
	//make random number
	rand.Seed(time.Now().UTC().UnixNano())
	randomWater := rand.Intn(100) //0 ~ 100
	randomWind := rand.Intn(100)  //0 ~ 100
	data := Data{}
	data.Water = int64(randomWater)
	data.Wind = int64(randomWind)
	status := Status{
		Status: data,
	}
	//write struct to json
	jsonData, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err.Error())
	}
	//update json file with new random number
	err = ioutil.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func OpenAndReadFile() (Status, error) {
	//open json file
	rootPath, err := os.Getwd()
	if err != nil {
		return Status{}, err
	}
	pathFile := path.Join(rootPath, "./data.json")
	file, err := os.OpenFile(pathFile, os.O_RDWR, 0644)
	if err != nil {
		return Status{}, err
	}

	defer file.Close()
	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return Status{}, err
	}
	var status Status
	err = json.Unmarshal(byteValue, &status)
	if err != nil {
		return Status{}, err
	}
	//return
	return status, nil
}
