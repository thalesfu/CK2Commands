package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2120] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2120][32120] = People_32120
	HistoryPeopleMap[2120][72120] = People_72120
	HistoryPeopleMap[2120][142120] = People_142120
	HistoryPeopleMap[2120][252120] = People_252120
	HistoryPeopleMap[2120][462120] = People_462120
}
