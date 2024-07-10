package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[999][30999] = People_30999
	HistoryPeopleMap[999][40999] = People_40999
	HistoryPeopleMap[999][70999] = People_70999
	HistoryPeopleMap[999][180999] = People_180999
	HistoryPeopleMap[999][260999] = People_260999
}
