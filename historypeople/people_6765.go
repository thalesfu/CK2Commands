package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6765] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6765][166765] = People_166765
	HistoryPeopleMap[6765][206765] = People_206765
}
