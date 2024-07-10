package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2097] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2097][12097] = People_12097
	HistoryPeopleMap[2097][72097] = People_72097
	HistoryPeopleMap[2097][142097] = People_142097
	HistoryPeopleMap[2097][252097] = People_252097
	HistoryPeopleMap[2097][462097] = People_462097
}
