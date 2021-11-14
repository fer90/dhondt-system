package dhondt

import "github.com/fer90/dhondt-system/party"

func Dhondt(parties []party.Party, availableSeats int) map[string]int {
	result := make(map[string]int)
	for _,party := range parties {
		result[party.Name] = availableSeats / cap(parties)
	}
	return result
}
