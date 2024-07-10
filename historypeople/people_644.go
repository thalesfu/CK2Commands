package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[644] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[644][644] = People_644
	HistoryPeopleMap[644][70644] = People_70644
	HistoryPeopleMap[644][260644] = People_260644
}
