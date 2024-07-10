package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[920] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[920][920] = People_920
	HistoryPeopleMap[920][30920] = People_30920
	HistoryPeopleMap[920][70920] = People_70920
	HistoryPeopleMap[920][260920] = People_260920
}
