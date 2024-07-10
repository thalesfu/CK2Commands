package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[123] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[123][123] = People_123
	HistoryPeopleMap[123][70123] = People_70123
	HistoryPeopleMap[123][90123] = People_90123
	HistoryPeopleMap[123][160123] = People_160123
	HistoryPeopleMap[123][170123] = People_170123
	HistoryPeopleMap[123][180123] = People_180123
	HistoryPeopleMap[123][190123] = People_190123
	HistoryPeopleMap[123][200123] = People_200123
	HistoryPeopleMap[123][470123] = People_470123
}
