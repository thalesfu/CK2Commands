package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3306] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3306][33306] = People_33306
	HistoryPeopleMap[3306][73306] = People_73306
	HistoryPeopleMap[3306][83306] = People_83306
	HistoryPeopleMap[3306][93306] = People_93306
}
