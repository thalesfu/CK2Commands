package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[543] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[543][543] = People_543
	HistoryPeopleMap[543][20543] = People_20543
	HistoryPeopleMap[543][30543] = People_30543
	HistoryPeopleMap[543][40543] = People_40543
	HistoryPeopleMap[543][70543] = People_70543
	HistoryPeopleMap[543][100543] = People_100543
	HistoryPeopleMap[543][190543] = People_190543
	HistoryPeopleMap[543][200543] = People_200543
	HistoryPeopleMap[543][260543] = People_260543
}
