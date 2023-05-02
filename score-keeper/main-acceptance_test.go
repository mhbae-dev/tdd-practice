package scoreKeeper

//void scoreTeamA1()
//void scoreTeamA2()
//void scoreTeamA3()
//void scoreTeamB1()
//void scoreTeamB2()
//void scoreTeamB3()
//String getScore()

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcceptancePrintScore(t *testing.T) {

	scoreBoard := &ScoreBoard{}
	scoreFormatter := &ScoreFormatter{}
	scoreKeeper := ScoreKeeper{scoreBoard, scoreFormatter}

	scoreKeeper.scoreTeamA1()
	scoreKeeper.scoreTeamA2()
	scoreKeeper.scoreTeamA3()
	scoreKeeper.scoreTeamB1()
	scoreKeeper.scoreTeamB2()
	scoreKeeper.scoreTeamB3()

	score := scoreKeeper.getScore()

	assert.Equal(t, "006:006", score)
}
