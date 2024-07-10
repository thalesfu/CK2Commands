package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1123] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1123][1123] = People_1123
	HistoryPeopleMap[1123][31123] = People_31123
	HistoryPeopleMap[1123][71123] = People_71123
	HistoryPeopleMap[1123][91123] = People_91123
	HistoryPeopleMap[1123][161123] = People_161123
	HistoryPeopleMap[1123][191123] = People_191123
	HistoryPeopleMap[1123][261123] = People_261123
}
