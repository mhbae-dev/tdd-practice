package scoreKeeper

import (
	"github.com/stretchr/testify/mock"
)

type MockPrinter struct {
	mock.Mock
}

func (mp MockPrinter) Print(scoreA int, scoreB int) string {
	args := mp.Called(scoreA, scoreB)
	return args.String(0)
}
