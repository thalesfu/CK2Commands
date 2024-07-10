package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1808] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1808][31808] = People_31808
	HistoryPeopleMap[1808][71808] = People_71808
	HistoryPeopleMap[1808][191808] = People_191808
}
