package shared

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var Tell func(text ...string)

func MainTell(text ...string) {
	builder := strings.Builder{}

	for _, t := range text {
		builder.WriteString(t)
	}

	fmt.Println(builder.String())
}

func Input() chan string {
	line := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}
	}()

	return line
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
