package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2900][32900] = People_32900
	HistoryPeopleMap[2900][72900] = People_72900
	HistoryPeopleMap[2900][142900] = People_142900
	HistoryPeopleMap[2900][472900] = People_472900
}
