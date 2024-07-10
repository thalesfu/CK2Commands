package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7776] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7776][7776] = People_7776
	HistoryPeopleMap[7776][167776] = People_167776
	HistoryPeopleMap[7776][247776] = People_247776
}
