package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1845] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1845][31845] = People_31845
	HistoryPeopleMap[1845][71845] = People_71845
	HistoryPeopleMap[1845][191845] = People_191845
}
