package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[960] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[960][960] = People_960
	HistoryPeopleMap[960][30960] = People_30960
	HistoryPeopleMap[960][70960] = People_70960
	HistoryPeopleMap[960][260960] = People_260960
}
