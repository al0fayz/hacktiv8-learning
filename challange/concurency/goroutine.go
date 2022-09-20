package main

import (
	"fmt"
)

type Message struct {
	Pesan string
}

func (p *Message) PrintMessage() {
	// jumlah := 3
	// message := "bisa"
	// var hasil [4]string
	// for i := 0; i < jumlah; i++ {
	// 	hasil[i] = message
	// }
	fmt.Println("bisa")
}
func (p *Message) PrintMessage2() {
	// jumlah := 3
	// message := "coba"
	// var hasil [4]string
	// for i := 0; i < jumlah; i++ {
	// 	hasil[i] = message
	// }
	fmt.Println("coba")
}

type example interface {
	PrintMessage()
	PrintMessage2()
}

func GoroutineAcak() {
	var ex example

	ex.PrintMessage()
	go ex.PrintMessage2()
}
