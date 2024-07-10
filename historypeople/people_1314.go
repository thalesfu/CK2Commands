package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1314] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1314][1314] = People_1314
	HistoryPeopleMap[1314][41314] = People_41314
	HistoryPeopleMap[1314][91314] = People_91314
	HistoryPeopleMap[1314][161314] = People_161314
	HistoryPeopleMap[1314][191314] = People_191314
	HistoryPeopleMap[1314][261314] = People_261314
}
