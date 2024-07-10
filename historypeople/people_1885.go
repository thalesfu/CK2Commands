package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1885] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1885][31885] = People_31885
	HistoryPeopleMap[1885][71885] = People_71885
	HistoryPeopleMap[1885][191885] = People_191885
}
