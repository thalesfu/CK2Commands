package lin

import (
	"CK2Commands/people"
)

func Hy() {
	people.Pollinate(getPeoples(), "lin")
}

func getPeoples() []people.Couple {
	var couples []people.Couple
	couples = append(couples, people.Couple{Husband: 2787625, Wife: 2793181}) // 林确 武燕
	return couples
}
