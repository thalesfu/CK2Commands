package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2303] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2303][12303] = People_12303
	HistoryPeopleMap[2303][32303] = People_32303
	HistoryPeopleMap[2303][72303] = People_72303
	HistoryPeopleMap[2303][82303] = People_82303
	HistoryPeopleMap[2303][142303] = People_142303
	HistoryPeopleMap[2303][242303] = People_242303
	HistoryPeopleMap[2303][252303] = People_252303
	HistoryPeopleMap[2303][462303] = People_462303
}
