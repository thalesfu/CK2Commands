package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[199] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[199][199] = People_199
	HistoryPeopleMap[199][160199] = People_160199
	HistoryPeopleMap[199][170199] = People_170199
	HistoryPeopleMap[199][190199] = People_190199
	HistoryPeopleMap[199][200199] = People_200199
}
