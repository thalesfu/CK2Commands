package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[970] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[970][30970] = People_30970
	HistoryPeopleMap[970][70970] = People_70970
	HistoryPeopleMap[970][260970] = People_260970
}
