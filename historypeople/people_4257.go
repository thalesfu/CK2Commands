package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4257] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4257][4257] = People_4257
	HistoryPeopleMap[4257][34257] = People_34257
	HistoryPeopleMap[4257][194257] = People_194257
}
