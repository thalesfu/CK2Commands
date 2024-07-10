package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7788] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7788][167788] = People_167788
	HistoryPeopleMap[7788][247788] = People_247788
}
