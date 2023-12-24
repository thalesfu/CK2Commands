package lou

import "github.com/thalesfu/CK2Commands/people"

func Pollinate() {
	people.Pollinate(getPollinatePeople(), "lou")
}

func getPollinatePeople() []people.Couple {
	var couples []people.Couple

	//couples = append(couples, people.Couple{Husband: M楼麟10961209, Wife: 2868714})

	return couples
}
