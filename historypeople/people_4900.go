package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4900][34900] = People_34900
	HistoryPeopleMap[4900][214900] = People_214900
}
