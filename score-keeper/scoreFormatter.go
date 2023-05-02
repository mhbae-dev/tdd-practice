package scoreKeeper

import "fmt"

type ScoreFormatter struct{}

type ScoreFormatterI interface {
	FormatScore(score []Score) string
}

func (sf *ScoreFormatter) FormatScore(scoreSlice []Score) string {
	score := make(map[string]int)
	for _, s := range scoreSlice {
		score[s.team] += s.points
	}
	return fmt.Sprintf("%03d:%03d", score["A"], score["B"])
}
