package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5995] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5995][5995] = People_5995
	HistoryPeopleMap[5995][125995] = People_125995
	HistoryPeopleMap[5995][145995] = People_145995
	HistoryPeopleMap[5995][205995] = People_205995
}
