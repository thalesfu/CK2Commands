package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1807] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1807][31807] = People_31807
	HistoryPeopleMap[1807][71807] = People_71807
	HistoryPeopleMap[1807][191807] = People_191807
}
