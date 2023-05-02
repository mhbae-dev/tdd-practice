package scoreKeeper

import (
	"reflect"
	"testing"
)

func TestScoreBoardShould(t *testing.T) {
	t.Run("AddScore stores score", func(t *testing.T) {
		scoreBoard := &ScoreBoard{}
		scoreBoard.AddScore(Score{"A", 1})
		want := Score{"A", 1}
		got := scoreBoard.scores[0]
		if got != want {
			t.Errorf("Expected scoreBoard.scores[0] to be Score{A, 1}, got %v", scoreBoard.scores[0])
		}
	})

	t.Run("GetScore returns scores", func(t *testing.T) {
		scoreBoard := &ScoreBoard{}
		scoreBoard.AddScore(Score{"A", 1})
		scoreBoard.AddScore(Score{"A", 2})
		want := []Score{{"A", 1}, {"A", 2}}
		got := scoreBoard.GetScore()
		if reflect.DeepEqual(got, want) == false {
			t.Errorf("Expected scoreBoard.GetScore() to be %v, got %v", want, got)
		}
	})
}
