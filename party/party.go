package party

type Party struct {
	Name string
	ValidVotes int
}

func New(name string, validVotes int) Party {
	newParty := Party {name, validVotes}
	return newParty
}
