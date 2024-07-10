package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2901] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2901][32901] = People_32901
	HistoryPeopleMap[2901][72901] = People_72901
	HistoryPeopleMap[2901][142901] = People_142901
	HistoryPeopleMap[2901][472901] = People_472901
}
