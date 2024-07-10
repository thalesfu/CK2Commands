package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1865] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1865][31865] = People_31865
	HistoryPeopleMap[1865][71865] = People_71865
	HistoryPeopleMap[1865][81865] = People_81865
	HistoryPeopleMap[1865][191865] = People_191865
}
