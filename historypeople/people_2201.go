package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2201] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2201][12201] = People_12201
	HistoryPeopleMap[2201][32201] = People_32201
	HistoryPeopleMap[2201][72201] = People_72201
	HistoryPeopleMap[2201][82201] = People_82201
	HistoryPeopleMap[2201][142201] = People_142201
	HistoryPeopleMap[2201][252201] = People_252201
}
