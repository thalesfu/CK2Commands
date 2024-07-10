package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[395] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[395][395] = People_395
	HistoryPeopleMap[395][20395] = People_20395
	HistoryPeopleMap[395][190395] = People_190395
	HistoryPeopleMap[395][260395] = People_260395
}
