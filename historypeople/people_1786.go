package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1786] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1786][31786] = People_31786
	HistoryPeopleMap[1786][191786] = People_191786
}
