package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2009] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2009][2009] = People_2009
	HistoryPeopleMap[2009][12009] = People_12009
	HistoryPeopleMap[2009][32009] = People_32009
	HistoryPeopleMap[2009][42009] = People_42009
	HistoryPeopleMap[2009][72009] = People_72009
	HistoryPeopleMap[2009][142009] = People_142009
	HistoryPeopleMap[2009][172009] = People_172009
	HistoryPeopleMap[2009][182009] = People_182009
	HistoryPeopleMap[2009][192009] = People_192009
	HistoryPeopleMap[2009][232009] = People_232009
	HistoryPeopleMap[2009][252009] = People_252009
	HistoryPeopleMap[2009][332009] = People_332009
}
