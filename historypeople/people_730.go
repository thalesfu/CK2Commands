package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[730] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[730][730] = People_730
	HistoryPeopleMap[730][20730] = People_20730
	HistoryPeopleMap[730][70730] = People_70730
	HistoryPeopleMap[730][260730] = People_260730
	HistoryPeopleMap[730][450730] = People_450730
}
