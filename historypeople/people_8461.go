package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8461] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8461][168461] = People_168461
	HistoryPeopleMap[8461][188461] = People_188461
}
