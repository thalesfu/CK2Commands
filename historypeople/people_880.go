package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[880] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[880][30880] = People_30880
	HistoryPeopleMap[880][70880] = People_70880
	HistoryPeopleMap[880][260880] = People_260880
}
