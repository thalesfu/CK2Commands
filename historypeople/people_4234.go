package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4234] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4234][34234] = People_34234
	HistoryPeopleMap[4234][144234] = People_144234
	HistoryPeopleMap[4234][194234] = People_194234
}
