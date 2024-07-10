package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4608] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4608][34608] = People_34608
	HistoryPeopleMap[4608][74608] = People_74608
	HistoryPeopleMap[4608][144608] = People_144608
	HistoryPeopleMap[4608][454608] = People_454608
}
