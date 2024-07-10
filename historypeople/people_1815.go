package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1815] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1815][71815] = People_71815
	HistoryPeopleMap[1815][81815] = People_81815
	HistoryPeopleMap[1815][191815] = People_191815
}
