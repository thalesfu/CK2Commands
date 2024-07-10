package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7839] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7839][7839] = People_7839
	HistoryPeopleMap[7839][167839] = People_167839
	HistoryPeopleMap[7839][247839] = People_247839
}
