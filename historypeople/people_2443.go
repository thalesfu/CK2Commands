package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2443] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2443][72443] = People_72443
	HistoryPeopleMap[2443][142443] = People_142443
	HistoryPeopleMap[2443][212443] = People_212443
}
