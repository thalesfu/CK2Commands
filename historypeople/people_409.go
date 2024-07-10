package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[409] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[409][409] = People_409
	HistoryPeopleMap[409][30409] = People_30409
	HistoryPeopleMap[409][40409] = People_40409
	HistoryPeopleMap[409][180409] = People_180409
	HistoryPeopleMap[409][190409] = People_190409
	HistoryPeopleMap[409][260409] = People_260409
}
