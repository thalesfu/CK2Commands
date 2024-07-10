package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2250] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2250][12250] = People_12250
	HistoryPeopleMap[2250][32250] = People_32250
	HistoryPeopleMap[2250][72250] = People_72250
	HistoryPeopleMap[2250][82250] = People_82250
	HistoryPeopleMap[2250][142250] = People_142250
	HistoryPeopleMap[2250][252250] = People_252250
	HistoryPeopleMap[2250][472250] = People_472250
}
