package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[901] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[901][901] = People_901
	HistoryPeopleMap[901][30901] = People_30901
	HistoryPeopleMap[901][40901] = People_40901
	HistoryPeopleMap[901][70901] = People_70901
	HistoryPeopleMap[901][260901] = People_260901
}
