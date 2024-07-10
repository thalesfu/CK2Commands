package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3567] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3567][33567] = People_33567
	HistoryPeopleMap[3567][73567] = People_73567
	HistoryPeopleMap[3567][83567] = People_83567
}
