package scoreKeeper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockScoreBoard struct {
	mock.Mock
}

func (b *MockScoreBoard) AddScore(score Score) {
	b.Called(score)
}

func (b *MockScoreBoard) GetScore() []Score {
	arg := b.Called()
	return arg.Get(0).([]Score)
}

type MockScoreFormatter struct {
	mock.Mock
}

func (sf *MockScoreFormatter) FormatScore(score []Score) string {
	arg := sf.Called(score)
	return arg.String(0)
}

func TestScoreKeeper(t *testing.T) {
	tests := []struct {
		name   string
		method func(sk *ScoreKeeper)
		score  Score
	}{
		{
			name:   "scoreTeamA1 stores score",
			method: func(sk *ScoreKeeper) { sk.scoreTeamA1() },
			score:  Score{team: "A", points: 1},
		},
		{
			name:   "scoreTeamA2 stores score",
			method: func(sk *ScoreKeeper) { sk.scoreTeamA2() },
			score:  Score{team: "A", points: 2},
		},
		{
			name:   "scoreTeamA3 stores score",
			method: func(sk *ScoreKeeper) { sk.scoreTeamA3() },
			score:  Score{team: "A", points: 3},
		},
		{
			name:   "scoreTeamB1 stores score",
			method: func(sk *ScoreKeeper) { sk.scoreTeamB1() },
			score:  Score{team: "B", points: 1},
		},
		{
			name:   "scoreTeamB2 stores score",
			method: func(sk *ScoreKeeper) { sk.scoreTeamB2() },
			score:  Score{team: "B", points: 2},
		},
		{
			name:   "scoreTeamB3 stores score",
			method: func(sk *ScoreKeeper) { sk.scoreTeamB3() },
			score:  Score{team: "B", points: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scoreBoard := new(MockScoreBoard)
			scoreFormatter := new(MockScoreFormatter)
			scoreBoard.On("AddScore", tt.score).Return()

			scoreKeeper := ScoreKeeper{scoreBoard, scoreFormatter}
			tt.method(&scoreKeeper)

			scoreBoard.AssertCalled(t, "AddScore", tt.score)
			scoreBoard.AssertExpectations(t)
		})
	}
}

func TestGetScore(t *testing.T) {
	scoreBoard := new(MockScoreBoard)
	scoreFormatter := new(MockScoreFormatter)
	scoreBoard.On("GetScore").Return(map[string]int{
		"A": 1,
		"B": 0,
	})

	// Set up the expected formatted score value on the mock ScoreFormatter.
	scoreFormatter.On("FormatScore", map[string]int{
		"A": 1,
		"B": 0,
	}).Return("001:000")

	scoreBoard.On("AddScore", Score{"A", 1}).Return()

	scoreKeeper := ScoreKeeper{scoreBoard, scoreFormatter}
	scoreKeeper.scoreTeamA1()

	score := scoreKeeper.getScore()

	scoreBoard.AssertCalled(t, "GetScore")
	scoreBoard.AssertCalled(t, "AddScore", Score{"A", 1})
	scoreFormatter.AssertCalled(t, "FormatScore", map[string]int{
		"A": 1,
		"B": 0,
	})

	scoreBoard.AssertExpectations(t)
	scoreFormatter.AssertExpectations(t)

	assert.Equal(t, "001:000", score)
}

//
//func TestScoreKeeperShould(t *testing.T) {
//	t.Run("scoreTeamA1 stores score", func(t *testing.T) {
//		scoreBoard := new(MockScoreBoard)
//		scoreBoard.On("AddScore", Score{
//			team:   "A",
//			points: 1,
//		}).Return()
//
//		scoreKeeper := ScoreKeeper{scoreBoard}
//		scoreKeeper.scoreTeamA1()
//
//		scoreBoard.AssertCalled(t, "AddScore", Score{"A", 1})
//		scoreBoard.AssertExpectations(t)
//	})
//
//	t.Run("scoreTeamA2 stores score", func(t *testing.T) {
//		scoreBoard := new(MockScoreBoard)
//		scoreBoard.On("AddScore", Score{
//			team:   "A",
//			points: 2,
//		}).Return()
//
//		scoreKeeper := ScoreKeeper{scoreBoard}
//		scoreKeeper.scoreTeamA2()
//
//		scoreBoard.AssertCalled(t, "AddScore", Score{"A", 2})
//		scoreBoard.AssertExpectations(t)
//	})
//}
