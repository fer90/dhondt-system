package dhondt

import (
	"errors"
	"strings"
	"strconv"
	"github.com/fer90/dhondt-system/party"
)

const MINIMAL_AVAILABLE_SEATS = 1

func Dhondt(parties []party.Party, availableSeats int) (map[string]int, error) {
	if (availableSeats < MINIMAL_AVAILABLE_SEATS) {
		return nil, errors.New("Available seats should be at least " + strconv.Itoa(MINIMAL_AVAILABLE_SEATS))
	}
	adjudicatedSeatByParty, votesRemaining := initializeDhondtCalculation(parties)
	for availableSeats > 0 {
		winnerPartyForNextSeat := calculatePartyForNextSeat(votesRemaining, parties)
		adjudicatedSeatByParty[winnerPartyForNextSeat.Name]++
		availableSeats--
		votesRemaining[winnerPartyForNextSeat.Name] = winnerPartyForNextSeat.ValidVotes / (adjudicatedSeatByParty[winnerPartyForNextSeat.Name] + 1)
	}
	return adjudicatedSeatByParty, nil
}

func initializeDhondtCalculation(parties []party.Party) (map[string]int, map[string]int) {
	result := make(map[string]int)
	votesRemaining := make(map[string]int)
	for _,party := range parties {
		result[party.Name] = 0
		votesRemaining[party.Name] = party.ValidVotes / (result[party.Name] + 1)
	}
	return result, votesRemaining
}

func calculatePartyForNextSeat(votesRemaining map[string]int, parties []party.Party) party.Party {
	partyNameWithMostVotes := getPartyNameWithMostVotes(votesRemaining)
	var winnerParty party.Party
	for _,party := range parties {
		if strings.Compare(partyNameWithMostVotes, party.Name) == 0 {
			winnerParty = party
		}
	}
	return winnerParty
}

func getPartyNameWithMostVotes(votesRemaining map[string]int) string {
	maxVotes := 0
	partyNameWithMostVotes := ""
	for partyName,votes := range votesRemaining {
		if votes > maxVotes {
			maxVotes = votes
			partyNameWithMostVotes = partyName
		}
	}
	return partyNameWithMostVotes
}