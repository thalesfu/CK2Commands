package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2627] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2627][32627] = People_32627
	HistoryPeopleMap[2627][72627] = People_72627
	HistoryPeopleMap[2627][142627] = People_142627
}
