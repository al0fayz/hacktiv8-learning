package assignment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type Biodata struct {
	Id        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func Assignment1() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("Tolong masukan nama yang di cari!")
		fmt.Println("Contoh: go run main.go Thomas")
	} else {
		//open and read file
		students, err := OpenAndReadFile()
		if err != nil {
			fmt.Println(err)
		}
		//check argument number or string
		if intArgs, err := strconv.Atoi(arg[1]); err == nil {
			//check by number or id
			for i := 0; i < len(students); i++ {
				if students[i].Id == intArgs {
					fmt.Println("ID :", students[i].Id)
					fmt.Println("Nama :", students[i].Nama)
					fmt.Println("Alamat :", students[i].Alamat)
					fmt.Println("Pekerjaan :", students[i].Pekerjaan)
					fmt.Println("Alasan :", students[i].Alasan)
					break
				}
			}
		} else {
			//search by nama
			for i := 0; i < len(students); i++ {
				if strings.Contains(students[i].Nama, arg[1]) {
					fmt.Println("ID :", students[i].Id)
					fmt.Println("Nama :", students[i].Nama)
					fmt.Println("Alamat :", students[i].Alamat)
					fmt.Println("Pekerjaan :", students[i].Pekerjaan)
					fmt.Println("Alasan :", students[i].Alasan)
					break
				}
			}
		}
	}
}

// function openfile dummy data
func OpenAndReadFile() ([]Biodata, error) {
	//open dummy file
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	pathFile := path.Join(rootPath, "/assignment/dummy.json")
	file, err := os.OpenFile(pathFile, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var students []Biodata
	err = json.Unmarshal(byteValue, &students)
	if err != nil {
		return nil, err
	}
	//return
	return students, nil
}
