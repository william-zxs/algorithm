package main

import (
	"fmt"
)

type AA struct{}

type BB struct {
	AA
}

func work(data AA) {
	fmt.Println("====")
}

func Play() {

}
