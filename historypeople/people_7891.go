package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7891] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7891][7891] = People_7891
	HistoryPeopleMap[7891][167891] = People_167891
	HistoryPeopleMap[7891][247891] = People_247891
}
