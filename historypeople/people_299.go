package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[299] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[299][299] = People_299
	HistoryPeopleMap[299][20299] = People_20299
	HistoryPeopleMap[299][70299] = People_70299
	HistoryPeopleMap[299][160299] = People_160299
	HistoryPeopleMap[299][170299] = People_170299
	HistoryPeopleMap[299][190299] = People_190299
	HistoryPeopleMap[299][200299] = People_200299
	HistoryPeopleMap[299][260299] = People_260299
}
