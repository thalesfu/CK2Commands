package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[895] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[895][30895] = People_30895
	HistoryPeopleMap[895][70895] = People_70895
	HistoryPeopleMap[895][260895] = People_260895
}
