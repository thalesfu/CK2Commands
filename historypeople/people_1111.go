package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1111] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1111][1111] = People_1111
	HistoryPeopleMap[1111][31111] = People_31111
	HistoryPeopleMap[1111][71111] = People_71111
	HistoryPeopleMap[1111][91111] = People_91111
	HistoryPeopleMap[1111][131111] = People_131111
	HistoryPeopleMap[1111][161111] = People_161111
	HistoryPeopleMap[1111][191111] = People_191111
	HistoryPeopleMap[1111][251111] = People_251111
	HistoryPeopleMap[1111][261111] = People_261111
}
