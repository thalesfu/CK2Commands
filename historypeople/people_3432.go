package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3432] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3432][33432] = People_33432
	HistoryPeopleMap[3432][83432] = People_83432
	HistoryPeopleMap[3432][93432] = People_93432
}
