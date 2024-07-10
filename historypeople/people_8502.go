package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8502] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8502][8502] = People_8502
	HistoryPeopleMap[8502][188502] = People_188502
	HistoryPeopleMap[8502][208502] = People_208502
	HistoryPeopleMap[8502][468502] = People_468502
	HistoryPeopleMap[8502][478502] = People_478502
}
