package scoreKeeper_test

import (
	scoreKeeper "github.com/mhbae-dev/tdd-practice/score-keeper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScorePrintShould(t *testing.T) {
	t.Run("pad score with zeros", func(t *testing.T) {
		scorePrinter := scoreKeeper.ScorePrinter{}
		score := scorePrinter.Print(2, 3)
		assert.Equal(t, "002:003", score)
	})

	t.Run("print score", func(t *testing.T) {
		scorePrinter := scoreKeeper.ScorePrinter{}
		score := scorePrinter.Print(6, 6)
		assert.Equal(t, "006:006", score)
	})
}
