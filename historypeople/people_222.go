package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[222] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[222][222] = People_222
	HistoryPeopleMap[222][40222] = People_40222
	HistoryPeopleMap[222][70222] = People_70222
	HistoryPeopleMap[222][170222] = People_170222
	HistoryPeopleMap[222][180222] = People_180222
	HistoryPeopleMap[222][190222] = People_190222
	HistoryPeopleMap[222][200222] = People_200222
	HistoryPeopleMap[222][260222] = People_260222
}
