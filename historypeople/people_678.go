package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[678] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[678][678] = People_678
	HistoryPeopleMap[678][20678] = People_20678
	HistoryPeopleMap[678][30678] = People_30678
	HistoryPeopleMap[678][70678] = People_70678
	HistoryPeopleMap[678][180678] = People_180678
	HistoryPeopleMap[678][260678] = People_260678
}
