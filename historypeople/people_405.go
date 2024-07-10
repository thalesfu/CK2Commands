package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[405] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[405][405] = People_405
	HistoryPeopleMap[405][30405] = People_30405
	HistoryPeopleMap[405][160405] = People_160405
	HistoryPeopleMap[405][180405] = People_180405
	HistoryPeopleMap[405][190405] = People_190405
	HistoryPeopleMap[405][260405] = People_260405
}
