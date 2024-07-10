package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[410] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[410][410] = People_410
	HistoryPeopleMap[410][30410] = People_30410
	HistoryPeopleMap[410][70410] = People_70410
	HistoryPeopleMap[410][180410] = People_180410
	HistoryPeopleMap[410][190410] = People_190410
	HistoryPeopleMap[410][260410] = People_260410
}
