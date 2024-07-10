package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1785] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1785][1785] = People_1785
	HistoryPeopleMap[1785][31785] = People_31785
	HistoryPeopleMap[1785][191785] = People_191785
}
