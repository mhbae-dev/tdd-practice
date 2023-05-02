package scoreKeeper_test

import (
	scoreKeeper "github.com/mhbae-dev/tdd-practice/score-keeper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreKeeperShould(t *testing.T) {
	t.Run("initial score is 000:000", func(t *testing.T) {
		mockPrinter := new(scoreKeeper.MockPrinter)
		scoreKeeper := scoreKeeper.NewScoreKeeper(mockPrinter)

		mockPrinter.On("Print", 0, 0).Return("000:000")

		want := "000:000"
		got := scoreKeeper.GetScore()

		assert.Equal(t, want, got)
		mockPrinter.AssertExpectations(t)
	})

	t.Run("score is 001:000 after team A scores 1 point", func(t *testing.T) {
		mockPrinter := new(scoreKeeper.MockPrinter)
		scoreKeeper := scoreKeeper.NewScoreKeeper(mockPrinter)

		mockPrinter.On("Print", 1, 0).Return("001:000")

		scoreKeeper.ScoreTeamA1()
		want := "001:000"
		got := scoreKeeper.GetScore()

		assert.Equal(t, want, got)
		mockPrinter.AssertExpectations(t)
	})

	t.Run("score is 002:000 after team A scores 2 points", func(t *testing.T) {
		mockPrinter := new(scoreKeeper.MockPrinter)
		scoreKeeper := scoreKeeper.NewScoreKeeper(mockPrinter)

		mockPrinter.On("Print", 2, 0).Return("002:000")

		scoreKeeper.ScoreTeamA2()
		want := "002:000"
		got := scoreKeeper.GetScore()

		assert.Equal(t, want, got)
		mockPrinter.AssertExpectations(t)
	})

	t.Run("score is 003:000 after team A scores 3 points", func(t *testing.T) {
		mockPrinter := new(scoreKeeper.MockPrinter)
		scoreKeeper := scoreKeeper.NewScoreKeeper(mockPrinter)

		mockPrinter.On("Print", 3, 0).Return("003:000")

		scoreKeeper.ScoreTeamA3()
		want := "003:000"
		got := scoreKeeper.GetScore()

		assert.Equal(t, want, got)
		mockPrinter.AssertExpectations(t)
	})

	t.Run("score is 000:001 after team B scores 1 point", func(t *testing.T) {
		mockPrinter := new(scoreKeeper.MockPrinter)
		scoreKeeper := scoreKeeper.NewScoreKeeper(mockPrinter)

		mockPrinter.On("Print", 0, 1).Return("000:001")

		scoreKeeper.ScoreTeamB1()
		want := "000:001"
		got := scoreKeeper.GetScore()

		assert.Equal(t, want, got)
		mockPrinter.AssertExpectations(t)
	})

	t.Run("score is 000:002 after team B scores 2 points", func(t *testing.T) {
		mockPrinter := new(scoreKeeper.MockPrinter)
		scoreKeeper := scoreKeeper.NewScoreKeeper(mockPrinter)

		mockPrinter.On("Print", 0, 2).Return("000:002")

		scoreKeeper.ScoreTeamB2()
		want := "000:002"
		got := scoreKeeper.GetScore()

		assert.Equal(t, want, got)
		mockPrinter.AssertExpectations(t)
	})

	t.Run("score is 000:003 after team B scores 3 points", func(t *testing.T) {
		mockPrinter := new(scoreKeeper.MockPrinter)
		scoreKeeper := scoreKeeper.NewScoreKeeper(mockPrinter)

		mockPrinter.On("Print", 0, 3).Return("000:003")

		scoreKeeper.ScoreTeamB3()
		want := "000:003"
		got := scoreKeeper.GetScore()

		assert.Equal(t, want, got)
		mockPrinter.AssertExpectations(t)
	})
}
