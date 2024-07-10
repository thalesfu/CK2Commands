package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4375] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4375][4375] = People_4375
	HistoryPeopleMap[4375][34375] = People_34375
	HistoryPeopleMap[4375][194375] = People_194375
}
