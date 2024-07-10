package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[127] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[127][127] = People_127
	HistoryPeopleMap[127][70127] = People_70127
	HistoryPeopleMap[127][90127] = People_90127
	HistoryPeopleMap[127][160127] = People_160127
	HistoryPeopleMap[127][170127] = People_170127
	HistoryPeopleMap[127][180127] = People_180127
	HistoryPeopleMap[127][190127] = People_190127
	HistoryPeopleMap[127][200127] = People_200127
	HistoryPeopleMap[127][470127] = People_470127
}
