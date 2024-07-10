package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1993] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1993][71993] = People_71993
	HistoryPeopleMap[1993][191993] = People_191993
	HistoryPeopleMap[1993][451993] = People_451993
}
