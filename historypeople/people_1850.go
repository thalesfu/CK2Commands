package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1850] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1850][31850] = People_31850
	HistoryPeopleMap[1850][71850] = People_71850
	HistoryPeopleMap[1850][81850] = People_81850
	HistoryPeopleMap[1850][191850] = People_191850
}
