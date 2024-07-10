package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2090] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2090][12090] = People_12090
	HistoryPeopleMap[2090][72090] = People_72090
	HistoryPeopleMap[2090][142090] = People_142090
	HistoryPeopleMap[2090][252090] = People_252090
	HistoryPeopleMap[2090][462090] = People_462090
}
