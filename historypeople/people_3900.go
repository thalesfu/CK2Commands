package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3900][3900] = People_3900
	HistoryPeopleMap[3900][33900] = People_33900
	HistoryPeopleMap[3900][73900] = People_73900
	HistoryPeopleMap[3900][183900] = People_183900
}
