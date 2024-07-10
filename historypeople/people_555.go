package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[555] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[555][30555] = People_30555
	HistoryPeopleMap[555][40555] = People_40555
	HistoryPeopleMap[555][70555] = People_70555
	HistoryPeopleMap[555][190555] = People_190555
	HistoryPeopleMap[555][260555] = People_260555
}
