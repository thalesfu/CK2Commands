package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1844] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1844][71844] = People_71844
	HistoryPeopleMap[1844][191844] = People_191844
}
