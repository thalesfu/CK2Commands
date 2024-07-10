package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1995] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1995][71995] = People_71995
	HistoryPeopleMap[1995][191995] = People_191995
	HistoryPeopleMap[1995][211995] = People_211995
	HistoryPeopleMap[1995][451995] = People_451995
}
