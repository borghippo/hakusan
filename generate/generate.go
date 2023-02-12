package generate

type HiraganaCharacter struct {
	Hiragana rune
	Latin    string
}

func GenerateHiragana() []HiraganaCharacter {
	return []HiraganaCharacter{{'あ', "a"}, {'い', "i"}, {'う', "u"}, {'え', "e"}, {'お', "o"}}
}
