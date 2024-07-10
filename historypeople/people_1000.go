package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1000] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1000][1000] = People_1000
	HistoryPeopleMap[1000][11000] = People_11000
	HistoryPeopleMap[1000][31000] = People_31000
	HistoryPeopleMap[1000][41000] = People_41000
	HistoryPeopleMap[1000][71000] = People_71000
	HistoryPeopleMap[1000][91000] = People_91000
	HistoryPeopleMap[1000][151000] = People_151000
	HistoryPeopleMap[1000][161000] = People_161000
	HistoryPeopleMap[1000][191000] = People_191000
	HistoryPeopleMap[1000][221000] = People_221000
	HistoryPeopleMap[1000][231000] = People_231000
	HistoryPeopleMap[1000][251000] = People_251000
	HistoryPeopleMap[1000][261000] = People_261000
	HistoryPeopleMap[1000][451000] = People_451000
}
