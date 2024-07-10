package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8251] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8251][168251] = People_168251
	HistoryPeopleMap[8251][188251] = People_188251
}
