package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9604] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9604][109604] = People_109604
	HistoryPeopleMap[9604][159604] = People_159604
}
