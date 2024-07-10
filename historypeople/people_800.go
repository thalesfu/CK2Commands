package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[800][20800] = People_20800
	HistoryPeopleMap[800][30800] = People_30800
	HistoryPeopleMap[800][40800] = People_40800
	HistoryPeopleMap[800][70800] = People_70800
	HistoryPeopleMap[800][220800] = People_220800
	HistoryPeopleMap[800][250800] = People_250800
	HistoryPeopleMap[800][450800] = People_450800
}
