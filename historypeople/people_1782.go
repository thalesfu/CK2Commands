package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1782] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1782][1782] = People_1782
	HistoryPeopleMap[1782][31782] = People_31782
	HistoryPeopleMap[1782][191782] = People_191782
}
