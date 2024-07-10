package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1828] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1828][71828] = People_71828
	HistoryPeopleMap[1828][81828] = People_81828
	HistoryPeopleMap[1828][191828] = People_191828
}
