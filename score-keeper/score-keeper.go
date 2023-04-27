package scoreKeeper

type ScoreKeeper struct {
	scorePrinter Printer
	teamAScore   int
	teamBScore   int
}

func NewScoreKeeper(scorePrinter Printer) *ScoreKeeper {
	return &ScoreKeeper{
		scorePrinter: scorePrinter,
	}
}

func (sk *ScoreKeeper) GetScore() string {
	return sk.scorePrinter.Print(sk.teamAScore, sk.teamBScore)
}

func (sk *ScoreKeeper) ScoreTeamA1() {
	sk.teamAScore += 1
}

func (sk *ScoreKeeper) ScoreTeamA2() {
	sk.teamAScore += 2
}

func (sk *ScoreKeeper) ScoreTeamA3() {
	sk.teamAScore += 3
}

func (sk *ScoreKeeper) ScoreTeamB1() {
	sk.teamBScore += 1
}

func (sk *ScoreKeeper) ScoreTeamB2() {
	sk.teamBScore += 2
}

func (sk *ScoreKeeper) ScoreTeamB3() {
	sk.teamBScore += 3
}
