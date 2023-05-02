package scoreKeeper_test

import (
	scoreKeeper "github.com/mhbae-dev/tdd-practice/score-keeper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreKeeperAcceptance(t *testing.T) {
	printer := scoreKeeper.ScorePrinter{}
	scoreKeeper := scoreKeeper.NewScoreKeeper(&printer)

	//mockPrinter.On("Print", 6, 6).Return("006:006")

	scoreKeeper.ScoreTeamA1()
	scoreKeeper.ScoreTeamA2()
	scoreKeeper.ScoreTeamA3()
	scoreKeeper.ScoreTeamB1()
	scoreKeeper.ScoreTeamB2()
	scoreKeeper.ScoreTeamB3()

	score := scoreKeeper.GetScore()
	assert.Equal(t, "006:006", score)
	//mockPrinter.AssertExpectations(t)
}
