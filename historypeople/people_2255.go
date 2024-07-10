package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2255] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2255][12255] = People_12255
	HistoryPeopleMap[2255][32255] = People_32255
	HistoryPeopleMap[2255][72255] = People_72255
	HistoryPeopleMap[2255][82255] = People_82255
	HistoryPeopleMap[2255][142255] = People_142255
	HistoryPeopleMap[2255][252255] = People_252255
}
