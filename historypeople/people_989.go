package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[989] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[989][30989] = People_30989
	HistoryPeopleMap[989][70989] = People_70989
	HistoryPeopleMap[989][260989] = People_260989
}
