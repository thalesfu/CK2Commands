package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2300] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2300][12300] = People_12300
	HistoryPeopleMap[2300][32300] = People_32300
	HistoryPeopleMap[2300][72300] = People_72300
	HistoryPeopleMap[2300][82300] = People_82300
	HistoryPeopleMap[2300][142300] = People_142300
	HistoryPeopleMap[2300][242300] = People_242300
	HistoryPeopleMap[2300][252300] = People_252300
	HistoryPeopleMap[2300][462300] = People_462300
	HistoryPeopleMap[2300][472300] = People_472300
}
