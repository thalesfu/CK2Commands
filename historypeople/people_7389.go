package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7389] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7389][7389] = People_7389
	HistoryPeopleMap[7389][217389] = People_217389
	HistoryPeopleMap[7389][247389] = People_247389
}
