package generate

import (
	"math/rand"
)

type HiraganaCharacter struct {
	Hiragana string
	Latin    string
}

func GenerateHiragana(options int) (HiraganaCharacter, []HiraganaCharacter) {
	hiraganaCharacters := []HiraganaCharacter{{"あ", "a"}, {"い", "i"}, {"う", "u"}, {"え", "e"}, {"お", "o"}}

	if options > len(hiraganaCharacters) {
		options = len(hiraganaCharacters)
	}

	rand.Shuffle(len(hiraganaCharacters), func(i, j int) {
		hiraganaCharacters[i], hiraganaCharacters[j] = hiraganaCharacters[j], hiraganaCharacters[i]
	})

	choices := hiraganaCharacters[:options]

	return choices[rand.Intn(len(choices))], choices

}
