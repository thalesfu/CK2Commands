package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[128] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[128][128] = People_128
	HistoryPeopleMap[128][40128] = People_40128
	HistoryPeopleMap[128][70128] = People_70128
	HistoryPeopleMap[128][90128] = People_90128
	HistoryPeopleMap[128][160128] = People_160128
	HistoryPeopleMap[128][170128] = People_170128
	HistoryPeopleMap[128][180128] = People_180128
	HistoryPeopleMap[128][190128] = People_190128
	HistoryPeopleMap[128][200128] = People_200128
	HistoryPeopleMap[128][470128] = People_470128
}
