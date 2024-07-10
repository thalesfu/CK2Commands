package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1788] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1788][31788] = People_31788
	HistoryPeopleMap[1788][191788] = People_191788
}
