package wu

import "github.com/thalesfu/CK2Commands/people"

func Pollinate() {
	people.Pollinate(getPollinatePeople(), "wu")
}

func getPollinatePeople() []people.Couple {
	var couples []people.Couple

	//couples = append(couples, people.Couple{Husband: 2864881, Wife: 2851908})  // 元超 武蕊

	return couples
}
