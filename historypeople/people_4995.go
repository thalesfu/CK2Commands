package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4995] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4995][4995] = People_4995
	HistoryPeopleMap[4995][144995] = People_144995
	HistoryPeopleMap[4995][454995] = People_454995
}
