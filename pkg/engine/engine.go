package engine

func Engine() (chan string, chan string) {
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
