package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3200][3200] = People_3200
	HistoryPeopleMap[3200][33200] = People_33200
	HistoryPeopleMap[3200][73200] = People_73200
	HistoryPeopleMap[3200][83200] = People_83200
	HistoryPeopleMap[3200][93200] = People_93200
	HistoryPeopleMap[3200][183200] = People_183200
	HistoryPeopleMap[3200][463200] = People_463200
}
