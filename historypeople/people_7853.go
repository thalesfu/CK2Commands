package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7853] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7853][7853] = People_7853
	HistoryPeopleMap[7853][167853] = People_167853
	HistoryPeopleMap[7853][247853] = People_247853
}
