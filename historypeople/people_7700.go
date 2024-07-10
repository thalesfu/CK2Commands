package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7700][7700] = People_7700
	HistoryPeopleMap[7700][167700] = People_167700
	HistoryPeopleMap[7700][217700] = People_217700
	HistoryPeopleMap[7700][247700] = People_247700
}
