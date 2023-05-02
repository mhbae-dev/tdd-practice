package scoreKeeper

type ScoreBoard struct {
	scores []Score
}

type Score struct {
	team   string
	points int
}

type ScoreBoardI interface {
	AddScore(score Score)
	GetScore() []Score
}

func (sb *ScoreBoard) AddScore(score Score) {
	sb.scores = append(sb.scores, score)
}

func (sb *ScoreBoard) GetScore() []Score {
	return sb.scores
}
