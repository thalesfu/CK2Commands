package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4640] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4640][74640] = People_74640
	HistoryPeopleMap[4640][144640] = People_144640
}
