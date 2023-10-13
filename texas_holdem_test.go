package poker_test

import (
	"fmt"
	"testing"
	"time"

	poker "github.com/comerc/go-app-via-tests"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(5)
		duration := 1 * time.Second
		if poker.IsBuild {
			duration *= 60
		}
		cases := []poker.ScheduledAlert{
			{At: 0 * duration, Amount: 100},
			{At: 10 * duration, Amount: 200},
			{At: 20 * duration, Amount: 300},
			{At: 30 * duration, Amount: 400},
			{At: 40 * duration, Amount: 500},
			{At: 50 * duration, Amount: 600},
			{At: 60 * duration, Amount: 800},
			{At: 70 * duration, Amount: 1000},
			{At: 80 * duration, Amount: 2000},
			{At: 90 * duration, Amount: 4000},
			{At: 100 * duration, Amount: 8000},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(7)
		duration := 1 * time.Second
		if poker.IsBuild {
			duration *= 60
		}
		cases := []poker.ScheduledAlert{
			{At: 0 * duration, Amount: 100},
			{At: 12 * duration, Amount: 200},
			{At: 24 * duration, Amount: 300},
			{At: 36 * duration, Amount: 400},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})

}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"
	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}
			got := blindAlerter.Alerts[i]
			poker.AssertScheduledAlert(t, got, want)
		})
	}
}
