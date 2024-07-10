package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5905] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5905][125905] = People_125905
	HistoryPeopleMap[5905][145905] = People_145905
}
