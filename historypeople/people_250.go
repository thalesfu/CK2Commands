package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[250] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[250][250] = People_250
	HistoryPeopleMap[250][20250] = People_20250
	HistoryPeopleMap[250][30250] = People_30250
	HistoryPeopleMap[250][70250] = People_70250
	HistoryPeopleMap[250][160250] = People_160250
	HistoryPeopleMap[250][170250] = People_170250
	HistoryPeopleMap[250][180250] = People_180250
	HistoryPeopleMap[250][190250] = People_190250
	HistoryPeopleMap[250][200250] = People_200250
	HistoryPeopleMap[250][260250] = People_260250
	HistoryPeopleMap[250][480250] = People_480250
}
