package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[600][600] = People_600
	HistoryPeopleMap[600][30600] = People_30600
	HistoryPeopleMap[600][40600] = People_40600
	HistoryPeopleMap[600][70600] = People_70600
	HistoryPeopleMap[600][110600] = People_110600
	HistoryPeopleMap[600][200600] = People_200600
	HistoryPeopleMap[600][260600] = People_260600
	HistoryPeopleMap[600][450600] = People_450600
	HistoryPeopleMap[600][460600] = People_460600
}
