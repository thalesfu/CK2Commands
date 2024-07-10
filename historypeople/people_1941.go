package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1941] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1941][71941] = People_71941
	HistoryPeopleMap[1941][191941] = People_191941
	HistoryPeopleMap[1941][451941] = People_451941
}
