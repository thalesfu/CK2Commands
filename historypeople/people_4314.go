package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4314] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4314][34314] = People_34314
	HistoryPeopleMap[4314][194314] = People_194314
}
