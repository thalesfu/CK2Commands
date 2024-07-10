package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4399] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4399][4399] = People_4399
	HistoryPeopleMap[4399][34399] = People_34399
	HistoryPeopleMap[4399][194399] = People_194399
}
