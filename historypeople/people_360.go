package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[360] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[360][360] = People_360
	HistoryPeopleMap[360][20360] = People_20360
	HistoryPeopleMap[360][30360] = People_30360
	HistoryPeopleMap[360][40360] = People_40360
	HistoryPeopleMap[360][70360] = People_70360
	HistoryPeopleMap[360][190360] = People_190360
	HistoryPeopleMap[360][200360] = People_200360
	HistoryPeopleMap[360][260360] = People_260360
}
