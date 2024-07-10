package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[765] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[765][765] = People_765
	HistoryPeopleMap[765][20765] = People_20765
	HistoryPeopleMap[765][30765] = People_30765
	HistoryPeopleMap[765][70765] = People_70765
	HistoryPeopleMap[765][180765] = People_180765
	HistoryPeopleMap[765][260765] = People_260765
}
