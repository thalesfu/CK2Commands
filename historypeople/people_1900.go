package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1900] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1900][71900] = People_71900
	HistoryPeopleMap[1900][161900] = People_161900
	HistoryPeopleMap[1900][191900] = People_191900
	HistoryPeopleMap[1900][461900] = People_461900
}
