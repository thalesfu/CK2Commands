package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[768] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[768][20768] = People_20768
	HistoryPeopleMap[768][30768] = People_30768
	HistoryPeopleMap[768][70768] = People_70768
	HistoryPeopleMap[768][260768] = People_260768
}
