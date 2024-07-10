package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[480] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[480][480] = People_480
	HistoryPeopleMap[480][20480] = People_20480
	HistoryPeopleMap[480][30480] = People_30480
	HistoryPeopleMap[480][70480] = People_70480
	HistoryPeopleMap[480][190480] = People_190480
	HistoryPeopleMap[480][260480] = People_260480
}
