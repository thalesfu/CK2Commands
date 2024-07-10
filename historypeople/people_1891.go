package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1891] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1891][31891] = People_31891
	HistoryPeopleMap[1891][71891] = People_71891
	HistoryPeopleMap[1891][191891] = People_191891
}
