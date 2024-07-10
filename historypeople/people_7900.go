package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7900][7900] = People_7900
	HistoryPeopleMap[7900][167900] = People_167900
	HistoryPeopleMap[7900][247900] = People_247900
}
