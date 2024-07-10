package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[914] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[914][914] = People_914
	HistoryPeopleMap[914][30914] = People_30914
	HistoryPeopleMap[914][70914] = People_70914
	HistoryPeopleMap[914][260914] = People_260914
}
