package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[900][900] = People_900
	HistoryPeopleMap[900][20900] = People_20900
	HistoryPeopleMap[900][30900] = People_30900
	HistoryPeopleMap[900][40900] = People_40900
	HistoryPeopleMap[900][70900] = People_70900
	HistoryPeopleMap[900][220900] = People_220900
	HistoryPeopleMap[900][260900] = People_260900
}
