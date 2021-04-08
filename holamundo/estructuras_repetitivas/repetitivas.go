package main

import "fmt"

func main() {
	curso := "aprende go"
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	//for range

	for _, valor := range curso {
		fmt.Println(string(valor))
	}
}
