package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2212] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2212][32212] = People_32212
	HistoryPeopleMap[2212][72212] = People_72212
	HistoryPeopleMap[2212][82212] = People_82212
	HistoryPeopleMap[2212][142212] = People_142212
	HistoryPeopleMap[2212][252212] = People_252212
}
