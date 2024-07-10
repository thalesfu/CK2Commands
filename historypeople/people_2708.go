package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2708] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2708][32708] = People_32708
	HistoryPeopleMap[2708][72708] = People_72708
	HistoryPeopleMap[2708][142708] = People_142708
}
