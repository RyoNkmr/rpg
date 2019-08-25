package entity

type Command struct {
	Text          string
	SecondaryText string
	ShortCutKey   rune
	Callback      func()
}
