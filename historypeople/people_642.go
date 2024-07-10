package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[642] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[642][642] = People_642
	HistoryPeopleMap[642][20642] = People_20642
	HistoryPeopleMap[642][30642] = People_30642
	HistoryPeopleMap[642][70642] = People_70642
	HistoryPeopleMap[642][110642] = People_110642
	HistoryPeopleMap[642][260642] = People_260642
}
