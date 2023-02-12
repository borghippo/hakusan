package generate

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type HiraganaCharacter struct {
	Hiragana string
	Latin    string
}

func GenerateHiragana(options int) (HiraganaCharacter, []HiraganaCharacter) {
	hiraganaCharacters := []HiraganaCharacter{{"あ", "a"}, {"い", "i"}, {"う", "u"}, {"え", "e"}, {"お", "o"}}
	choices := []HiraganaCharacter{}

	rand.Shuffle(len(hiraganaCharacters), func(i, j int) {
		hiraganaCharacters[i], hiraganaCharacters[j] = hiraganaCharacters[j], hiraganaCharacters[i]
	})

	for i := 0; i < options; i++ {
		choices = append(choices, hiraganaCharacters[i])
	}

	return choices[rand.Intn(len(choices))], choices

}
