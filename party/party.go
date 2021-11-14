package party

type Party struct {
	Name string
	validVotes int
}

func New(name string, validVotes int) Party {
	newParty := Party {name, validVotes}
	return newParty
}
