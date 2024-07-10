package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[380] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[380][380] = People_380
	HistoryPeopleMap[380][20380] = People_20380
	HistoryPeopleMap[380][30380] = People_30380
	HistoryPeopleMap[380][70380] = People_70380
	HistoryPeopleMap[380][160380] = People_160380
	HistoryPeopleMap[380][190380] = People_190380
	HistoryPeopleMap[380][200380] = People_200380
	HistoryPeopleMap[380][260380] = People_260380
}
