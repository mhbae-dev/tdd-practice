package scoreKeeper

type Printer interface {
	Print(scoreA int, scoreB int) string
}
