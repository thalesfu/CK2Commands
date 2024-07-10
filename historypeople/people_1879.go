package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1879] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1879][31879] = People_31879
	HistoryPeopleMap[1879][71879] = People_71879
	HistoryPeopleMap[1879][191879] = People_191879
}
