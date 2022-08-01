package main

import "fmt"

func modify(log *string) {

	*log += "asd"
}

func main() {
	log := "a"

	fmt.Println(log)
	modify(&log)
	fmt.Println(log)


}