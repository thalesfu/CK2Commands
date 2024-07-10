package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2425] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2425][12425] = People_12425
	HistoryPeopleMap[2425][72425] = People_72425
	HistoryPeopleMap[2425][142425] = People_142425
}
