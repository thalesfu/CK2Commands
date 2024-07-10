package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2480] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2480][72480] = People_72480
	HistoryPeopleMap[2480][142480] = People_142480
}
