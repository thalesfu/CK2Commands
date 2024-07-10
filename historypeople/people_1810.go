package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1810] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1810][31810] = People_31810
	HistoryPeopleMap[1810][71810] = People_71810
	HistoryPeopleMap[1810][191810] = People_191810
}
