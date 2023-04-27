package scoreKeeper

import "fmt"

type ScorePrinter struct {
}

func (sp *ScorePrinter) Print(teamAScore int, teamBScore int) string {
	return fmt.Sprintf("%03d:%03d", teamAScore, teamBScore)
}
