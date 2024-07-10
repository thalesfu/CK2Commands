package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7850] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7850][7850] = People_7850
	HistoryPeopleMap[7850][167850] = People_167850
	HistoryPeopleMap[7850][247850] = People_247850
}
