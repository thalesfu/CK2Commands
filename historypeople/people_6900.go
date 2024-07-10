package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6900][6900] = People_6900
	HistoryPeopleMap[6900][166900] = People_166900
	HistoryPeopleMap[6900][206900] = People_206900
}
