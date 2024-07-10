package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4492] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4492][4492] = People_4492
	HistoryPeopleMap[4492][34492] = People_34492
	HistoryPeopleMap[4492][74492] = People_74492
}
