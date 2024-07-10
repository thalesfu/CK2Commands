package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5123] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5123][125123] = People_125123
	HistoryPeopleMap[5123][145123] = People_145123
	HistoryPeopleMap[5123][155123] = People_155123
	HistoryPeopleMap[5123][205123] = People_205123
	HistoryPeopleMap[5123][235123] = People_235123
}
