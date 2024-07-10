package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1943] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1943][71943] = People_71943
	HistoryPeopleMap[1943][161943] = People_161943
	HistoryPeopleMap[1943][191943] = People_191943
}
