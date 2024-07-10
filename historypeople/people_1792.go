package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1792] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1792][1792] = People_1792
	HistoryPeopleMap[1792][31792] = People_31792
	HistoryPeopleMap[1792][191792] = People_191792
}
