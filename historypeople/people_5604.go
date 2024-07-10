package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5604] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5604][5604] = People_5604
	HistoryPeopleMap[5604][125604] = People_125604
	HistoryPeopleMap[5604][145604] = People_145604
	HistoryPeopleMap[5604][205604] = People_205604
}
