package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[400][400] = People_400
	HistoryPeopleMap[400][30400] = People_30400
	HistoryPeopleMap[400][40400] = People_40400
	HistoryPeopleMap[400][70400] = People_70400
	HistoryPeopleMap[400][190400] = People_190400
	HistoryPeopleMap[400][260400] = People_260400
}
