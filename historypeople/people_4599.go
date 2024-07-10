package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4599] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4599][34599] = People_34599
	HistoryPeopleMap[4599][74599] = People_74599
	HistoryPeopleMap[4599][144599] = People_144599
}
