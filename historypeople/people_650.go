package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[650] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[650][650] = People_650
	HistoryPeopleMap[650][20650] = People_20650
	HistoryPeopleMap[650][30650] = People_30650
	HistoryPeopleMap[650][70650] = People_70650
	HistoryPeopleMap[650][110650] = People_110650
	HistoryPeopleMap[650][260650] = People_260650
	HistoryPeopleMap[650][450650] = People_450650
}
