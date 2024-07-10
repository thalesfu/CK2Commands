package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4323] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4323][34323] = People_34323
	HistoryPeopleMap[4323][194323] = People_194323
}
