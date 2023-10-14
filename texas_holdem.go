package poker

import (
	"io"
	"time"

	utils "github.com/comerc/go-app-via-tests/internal/utils"
)

var IsBuild = utils.GetWorkDir() != "."

// TexasHoldem manages a game of poker.
type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

// NewTexasHoldem returns a new game.
func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

// Start will schedule blind alerts dependant on the number of players.
func (p *TexasHoldem) Start(numberOfPlayers int, to io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Second
	if IsBuild {
		blindIncrement *= 60
	}
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, to)
		blindTime = blindTime + blindIncrement
	}
}

// Finish ends the game, recording the winner.
func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}
