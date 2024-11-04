package main

import (
	"fmt"
	"strings"
)

func compressedString(word string) string {
	var output strings.Builder
	compress(word, &output)
	return output.String()
}

func compress(word string, output *strings.Builder) {
	if len(word) == 0 {
		return
	}

	selected := word[0]
	counter := 1

	for ; counter < len(word); counter++ {
		if word[counter] != selected || counter > 8 {
			break
		}
	}

	output.WriteString(fmt.Sprintf("%d%c", counter, selected))

	compress(word[counter:], output)
}
