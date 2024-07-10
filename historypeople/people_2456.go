package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2456] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2456][72456] = People_72456
	HistoryPeopleMap[2456][142456] = People_142456
}
