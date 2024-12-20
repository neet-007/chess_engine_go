package engine

import "fmt"

func Engine() (chan string, chan string) {
	fmt.Println("from engine")
	frEng := make(chan string)
	toEng := make(chan string)

	go func() {
		for cmd := range toEng {
			switch cmd {
			case "quit":
				{

				}
			case "stop":
				{

				}
			}
		}

	}()

	return frEng, toEng
}
