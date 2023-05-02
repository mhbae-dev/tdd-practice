package scoreKeeper

import "testing"

func TestScoreFormatter_FormatScore(t *testing.T) {
	scoreFormatter := &ScoreFormatter{}

	got := scoreFormatter.FormatScore([]Score{{"A", 1}, {"B", 2}})
	want := "001:002"

	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}
