package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2462] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2462][72462] = People_72462
	HistoryPeopleMap[2462][142462] = People_142462
}
