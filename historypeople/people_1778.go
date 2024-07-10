package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1778] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1778][1778] = People_1778
	HistoryPeopleMap[1778][31778] = People_31778
	HistoryPeopleMap[1778][191778] = People_191778
}
