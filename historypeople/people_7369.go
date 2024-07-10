package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7369] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7369][7369] = People_7369
	HistoryPeopleMap[7369][167369] = People_167369
	HistoryPeopleMap[7369][217369] = People_217369
	HistoryPeopleMap[7369][247369] = People_247369
}
