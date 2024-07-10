package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[192] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[192][192] = People_192
	HistoryPeopleMap[192][160192] = People_160192
	HistoryPeopleMap[192][170192] = People_170192
	HistoryPeopleMap[192][180192] = People_180192
	HistoryPeopleMap[192][190192] = People_190192
	HistoryPeopleMap[192][200192] = People_200192
	HistoryPeopleMap[192][470192] = People_470192
}
