package dhondt

import (
	"testing"
	"github.com/fer90/dhondt-system/party"
)

func TestAvailableSeatsShouldBeGreaterThanZero(t *testing.T) {
	const WRONG_AVAILABLE_SEATS = -1
	const PARTY_VALID_VOTES = 1000
	testParty := party.New ("Democrats", PARTY_VALID_VOTES)
	result, err := Dhondt([]party.Party{testParty}, WRONG_AVAILABLE_SEATS)
	if result != nil && err == nil {
		t.Errorf("Available seats should be greater than zero!")
	}
}

func TestOneParty(t *testing.T) {
	const AVAILABLE_SEATS = 2
	const PARTY_VALID_VOTES = 1000
	testParty := party.New ("Democrats", PARTY_VALID_VOTES)
	result, err := Dhondt([]party.Party{testParty}, AVAILABLE_SEATS)
	if err == nil && result["Democrats"] != AVAILABLE_SEATS {
		t.Errorf("The only party should get all the seats!")
	}
}

func TestTwoPartiesWithSameVotes(t *testing.T) {
	const AVAILABLE_SEATS = 2
	const SEATS_FOR_EACH_PARTY = 1
	const PARTY_VALID_VOTES = 1000
	testDemocratParty := party.New ("Democrats", PARTY_VALID_VOTES)
	testRepublicanParty := party.New ("Republicans", PARTY_VALID_VOTES)
	result, err := Dhondt([]party.Party{testDemocratParty, testRepublicanParty}, AVAILABLE_SEATS)
	if err == nil && result["Democrats"] != SEATS_FOR_EACH_PARTY || result["Republicans"] != SEATS_FOR_EACH_PARTY {
		t.Errorf("Both parties should have the same seats!")
	}
}

func TestTwoPartiesWithDifferentVotesButSameResult(t *testing.T) {
	const AVAILABLE_SEATS = 2
	const SEATS_FOR_EACH_PARTY = 1
	const DEMOCRAT_PARTY_VALID_VOTES = 1000
	const REPUBLICAN_PARTY_VALID_VOTES = 750
	testDemocratParty := party.New ("Democrats", DEMOCRAT_PARTY_VALID_VOTES)
	testRepublicanParty := party.New ("Republicans", REPUBLICAN_PARTY_VALID_VOTES)
	result, err := Dhondt([]party.Party{testDemocratParty, testRepublicanParty}, AVAILABLE_SEATS)
	if err == nil && result["Democrats"] != SEATS_FOR_EACH_PARTY || result["Republicans"] != SEATS_FOR_EACH_PARTY {
		t.Errorf("Both parties should have the same seats!")
	}
}

func TestTwoPartiesWithDifferentVotesAndOneGetsAllTheSeats(t *testing.T) {
	const AVAILABLE_SEATS = 2
	const DEMOCRAT_PARTY_VALID_VOTES = 1000
	const REPUBLICAN_PARTY_VALID_VOTES = 400
	testDemocratParty := party.New ("Democrats", DEMOCRAT_PARTY_VALID_VOTES)
	testRepublicanParty := party.New ("Republicans", REPUBLICAN_PARTY_VALID_VOTES)
	result, err := Dhondt([]party.Party{testDemocratParty, testRepublicanParty}, AVAILABLE_SEATS)
	if err == nil && result["Democrats"] != AVAILABLE_SEATS {
		t.Errorf("Democrats should get all the seats!")
	}
	if err == nil && result["Republicans"] != 0 {
		t.Errorf("Republicans should not get any seat!")
	}
}
