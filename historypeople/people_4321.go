package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4321] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4321][34321] = People_34321
	HistoryPeopleMap[4321][194321] = People_194321
}
