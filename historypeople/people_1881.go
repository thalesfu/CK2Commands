package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1881] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1881][71881] = People_71881
	HistoryPeopleMap[1881][191881] = People_191881
}
