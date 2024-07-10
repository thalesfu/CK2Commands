package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1857] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1857][31857] = People_31857
	HistoryPeopleMap[1857][71857] = People_71857
	HistoryPeopleMap[1857][81857] = People_81857
	HistoryPeopleMap[1857][191857] = People_191857
}
