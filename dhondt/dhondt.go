package dhondt

import (
	"errors"
	"strings"
	"github.com/fer90/dhondt-system/party"
)

const MINIMAL_AVAILABLE_SEATS = 1

func Dhondt(parties []party.Party, availableSeats int) (map[string]int, error) {
	if (availableSeats < MINIMAL_AVAILABLE_SEATS) {
		return nil, errors.New("Available seats should be greater than zero")
	}
	result := make(map[string]int)
	votesRemaining := make(map[string]int)
	for _,party := range parties {
		result[party.Name] = 0
		votesRemaining[party.Name] = party.ValidVotes / (result[party.Name] + 1)
	}
	adjudicatedSeats := 0
	for adjudicatedSeats != availableSeats {
		winnerPartyForNextSeat := calculatePartyForNextSeat(votesRemaining, parties)
		result[winnerPartyForNextSeat.Name]++
		adjudicatedSeats++
		votesRemaining[winnerPartyForNextSeat.Name] = winnerPartyForNextSeat.ValidVotes / (result[winnerPartyForNextSeat.Name] + 1)
	}
	return result, nil
}

func calculatePartyForNextSeat(votesRemaining map[string]int, parties []party.Party) party.Party {
	maxVotes := 0
	partyNameWithMostVotes := ""
	for partyName,votes := range votesRemaining {
		if votes > maxVotes {
			maxVotes = votes
			partyNameWithMostVotes = partyName
		}
	}
	var winnerParty party.Party
	for _,party := range parties {
		if strings.Compare(partyNameWithMostVotes, party.Name) == 0 {
			winnerParty = party
		}
	}
	return winnerParty
}
