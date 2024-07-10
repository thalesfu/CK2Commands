package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[720] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[720][720] = People_720
	HistoryPeopleMap[720][30720] = People_30720
	HistoryPeopleMap[720][70720] = People_70720
	HistoryPeopleMap[720][260720] = People_260720
	HistoryPeopleMap[720][450720] = People_450720
}
