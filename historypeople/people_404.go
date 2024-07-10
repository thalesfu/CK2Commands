package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[404] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[404][404] = People_404
	HistoryPeopleMap[404][30404] = People_30404
	HistoryPeopleMap[404][40404] = People_40404
	HistoryPeopleMap[404][70404] = People_70404
	HistoryPeopleMap[404][160404] = People_160404
	HistoryPeopleMap[404][190404] = People_190404
	HistoryPeopleMap[404][260404] = People_260404
}
