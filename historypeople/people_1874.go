package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1874] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1874][31874] = People_31874
	HistoryPeopleMap[1874][71874] = People_71874
	HistoryPeopleMap[1874][191874] = People_191874
}
