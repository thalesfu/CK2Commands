package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5900][125900] = People_125900
	HistoryPeopleMap[5900][145900] = People_145900
}
