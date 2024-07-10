package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2223] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2223][32223] = People_32223
	HistoryPeopleMap[2223][72223] = People_72223
	HistoryPeopleMap[2223][82223] = People_82223
	HistoryPeopleMap[2223][142223] = People_142223
	HistoryPeopleMap[2223][252223] = People_252223
	HistoryPeopleMap[2223][262223] = People_262223
}
