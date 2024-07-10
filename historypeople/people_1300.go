package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1300] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1300][41300] = People_41300
	HistoryPeopleMap[1300][71300] = People_71300
	HistoryPeopleMap[1300][91300] = People_91300
	HistoryPeopleMap[1300][161300] = People_161300
	HistoryPeopleMap[1300][191300] = People_191300
	HistoryPeopleMap[1300][261300] = People_261300
}
