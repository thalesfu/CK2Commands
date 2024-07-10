package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1983] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1983][71983] = People_71983
	HistoryPeopleMap[1983][161983] = People_161983
	HistoryPeopleMap[1983][191983] = People_191983
	HistoryPeopleMap[1983][451983] = People_451983
}
