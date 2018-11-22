package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	for i := 0; i < 100; i++ {
		//fmt.Println("rand:", generateRand(4))
		fmt.Println("rand:", generateRand1(4))
	}

}

func generateRand(length int) string {
	bs := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	bytes := make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = bs[b%9]
	}
	return string(bytes)
}

func generateRand1(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)
}
