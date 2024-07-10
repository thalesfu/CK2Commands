package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[630] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[630][630] = People_630
	HistoryPeopleMap[630][20630] = People_20630
	HistoryPeopleMap[630][30630] = People_30630
	HistoryPeopleMap[630][70630] = People_70630
	HistoryPeopleMap[630][180630] = People_180630
	HistoryPeopleMap[630][260630] = People_260630
	HistoryPeopleMap[630][450630] = People_450630
}
