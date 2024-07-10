package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[100] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[100][100] = People_100
	HistoryPeopleMap[100][30100] = People_30100
	HistoryPeopleMap[100][40100] = People_40100
	HistoryPeopleMap[100][70100] = People_70100
	HistoryPeopleMap[100][160100] = People_160100
	HistoryPeopleMap[100][170100] = People_170100
	HistoryPeopleMap[100][180100] = People_180100
	HistoryPeopleMap[100][190100] = People_190100
	HistoryPeopleMap[100][200100] = People_200100
	HistoryPeopleMap[100][260100] = People_260100
	HistoryPeopleMap[100][470100] = People_470100
	HistoryPeopleMap[100][480100] = People_480100
}
