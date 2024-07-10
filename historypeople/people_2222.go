package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2222] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2222][32222] = People_32222
	HistoryPeopleMap[2222][72222] = People_72222
	HistoryPeopleMap[2222][82222] = People_82222
	HistoryPeopleMap[2222][142222] = People_142222
	HistoryPeopleMap[2222][252222] = People_252222
	HistoryPeopleMap[2222][262222] = People_262222
}
