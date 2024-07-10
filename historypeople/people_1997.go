package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1997] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1997][71997] = People_71997
	HistoryPeopleMap[1997][191997] = People_191997
	HistoryPeopleMap[1997][451997] = People_451997
}
