package bi

import "github.com/thalesfu/CK2Commands/people"

func Pollinate() {
	people.Pollinate(getPollinatePeople(), "bi")
}

func getPollinatePeople() []people.Couple {
	var couples []people.Couple

	couples = append(couples, people.Couple{Husband: M毕藩10770717, Wife: 2878014})

	return couples
}
