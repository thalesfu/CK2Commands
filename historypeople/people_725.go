package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[725] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[725][725] = People_725
	HistoryPeopleMap[725][30725] = People_30725
	HistoryPeopleMap[725][70725] = People_70725
	HistoryPeopleMap[725][260725] = People_260725
	HistoryPeopleMap[725][450725] = People_450725
}
