package main

import (
	"fmt"
	"github.com/borghippo/hakusan/generate"
)

func main() {
	hiraganaCharacters := generate.GenerateHiragana()

	for _, char := range hiraganaCharacters {
		fmt.Printf("hiragana: %c latin: %s\n", char.Hiragana, char.Latin)
	}
}
