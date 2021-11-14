package dhondt

import "github.com/fer90/dhondt-system/party"

func Dhondt(parties []party.Party, availableSeats int) map[string]int {
	result := map[string]int {
		parties[0].Name: availableSeats,
	}
	return result
}
