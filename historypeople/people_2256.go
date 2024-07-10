package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2256] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2256][12256] = People_12256
	HistoryPeopleMap[2256][32256] = People_32256
	HistoryPeopleMap[2256][72256] = People_72256
	HistoryPeopleMap[2256][82256] = People_82256
	HistoryPeopleMap[2256][142256] = People_142256
	HistoryPeopleMap[2256][252256] = People_252256
}
