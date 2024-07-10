package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[303] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[303][30303] = People_30303
	HistoryPeopleMap[303][40303] = People_40303
	HistoryPeopleMap[303][70303] = People_70303
	HistoryPeopleMap[303][160303] = People_160303
	HistoryPeopleMap[303][170303] = People_170303
	HistoryPeopleMap[303][190303] = People_190303
	HistoryPeopleMap[303][200303] = People_200303
	HistoryPeopleMap[303][260303] = People_260303
	HistoryPeopleMap[303][420303] = People_420303
}
