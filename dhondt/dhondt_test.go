package dhondt

import (
	"testing"
	"github.com/fer90/dhondt-system/party"
)

func TestOneParty(t *testing.T) {
	const AVAILABLE_SEATS = 2
	const PARTY_VALID_VOTES = 1000
	testParty := party.New ("Democrats", PARTY_VALID_VOTES)
	result := Dhondt([]party.Party{testParty}, AVAILABLE_SEATS)
	if result["Democrats"] != AVAILABLE_SEATS {
		t.Errorf("The only party should get all the seats!")
	}
}

func TestTwoParties(t *testing.T) {
	const AVAILABLE_SEATS = 2
	const SEATS_FOR_EACH_PARTY = 1
	const PARTY_VALID_VOTES = 1000
	testDemocratParty := party.New ("Democrats", PARTY_VALID_VOTES)
	testRepublicanParty := party.New ("Republicans", PARTY_VALID_VOTES)
	result := Dhondt([]party.Party{testDemocratParty, testRepublicanParty}, AVAILABLE_SEATS)
	if result["Democrats"] != SEATS_FOR_EACH_PARTY || result["Republicans"] != SEATS_FOR_EACH_PARTY {
		t.Errorf("Both parties should have the same seats!")
	}
}
