package lin

import (
	"github.com/thalesfu/CK2Commands/people"
)

func Pollinate() {
	people.Pollinate(getPollinatePeople(), "lin")
}

func getPollinatePeople() []people.Couple {
	var couples []people.Couple
	couples = append(couples, people.Couple{Husband: M林绛10690610, Wife: 2868714})
	couples = append(couples, people.Couple{Husband: 2875976, Wife: F林畅10910523})
	couples = append(couples, people.Couple{Husband: M林承庆10960227, Wife: 2876245})
	couples = append(couples, people.Couple{Husband: 2876925, Wife: F林明月10960614})
	couples = append(couples, people.Couple{Husband: M林琯10961020, Wife: 2876934})
	couples = append(couples, people.Couple{Husband: 2877624, Wife: F林嫩娘10980327})
	return couples
}
