package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4567] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4567][34567] = People_34567
	HistoryPeopleMap[4567][74567] = People_74567
	HistoryPeopleMap[4567][144567] = People_144567
}
