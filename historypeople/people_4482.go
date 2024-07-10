package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4482] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4482][34482] = People_34482
	HistoryPeopleMap[4482][74482] = People_74482
}
