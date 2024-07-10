package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4999][34999] = People_34999
	HistoryPeopleMap[4999][144999] = People_144999
	HistoryPeopleMap[4999][394999] = People_394999
	HistoryPeopleMap[4999][454999] = People_454999
}
