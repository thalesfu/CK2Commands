package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4379] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4379][4379] = People_4379
	HistoryPeopleMap[4379][34379] = People_34379
	HistoryPeopleMap[4379][194379] = People_194379
}
