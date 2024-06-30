package main

import (
	"fmt"
	"time"
)

func fazerping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func fazerpong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Println("Desenvoldedor: Ramon Cintas Gomes")
	fmt.Println("A cada segundo é printado a palavra ping e depois pong")
	fmt.Println("Para parar o programa basta apertar a tecla ENTER")
	var c chan string = make(chan string)
	go fazerping(c)
	go imprimir(c)
	go fazerpong(c)
	var entrada string
	fmt.Scanln(&entrada)
}
