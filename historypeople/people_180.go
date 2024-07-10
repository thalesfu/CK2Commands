package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[180] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[180][180] = People_180
	HistoryPeopleMap[180][40180] = People_40180
	HistoryPeopleMap[180][70180] = People_70180
	HistoryPeopleMap[180][160180] = People_160180
	HistoryPeopleMap[180][170180] = People_170180
	HistoryPeopleMap[180][180180] = People_180180
	HistoryPeopleMap[180][190180] = People_190180
	HistoryPeopleMap[180][200180] = People_200180
	HistoryPeopleMap[180][470180] = People_470180
	HistoryPeopleMap[180][480180] = People_480180
}
