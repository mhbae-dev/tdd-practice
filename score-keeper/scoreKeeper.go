package scoreKeeper

type ScoreKeeper struct {
	scoreBoard     ScoreBoardI
	ScoreFormatter ScoreFormatterI
}

func (sk *ScoreKeeper) scoreTeamA1() {
	sk.scoreBoard.AddScore(Score{"A", 1})
}

func (sk *ScoreKeeper) scoreTeamA2() {
	sk.scoreBoard.AddScore(Score{"A", 2})
}

func (sk *ScoreKeeper) scoreTeamA3() {
	sk.scoreBoard.AddScore(Score{"A", 3})
}

func (sk *ScoreKeeper) scoreTeamB1() {
	sk.scoreBoard.AddScore(Score{"B", 1})
}

func (sk *ScoreKeeper) scoreTeamB2() {
	sk.scoreBoard.AddScore(Score{"B", 2})
}

func (sk *ScoreKeeper) scoreTeamB3() {
	sk.scoreBoard.AddScore(Score{"B", 3})
}

func (sk *ScoreKeeper) getScore() string {
	return sk.ScoreFormatter.FormatScore(sk.scoreBoard.GetScore())
}
