package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2380] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2380][12380] = People_12380
	HistoryPeopleMap[2380][72380] = People_72380
	HistoryPeopleMap[2380][142380] = People_142380
	HistoryPeopleMap[2380][212380] = People_212380
	HistoryPeopleMap[2380][252380] = People_252380
	HistoryPeopleMap[2380][462380] = People_462380
}
