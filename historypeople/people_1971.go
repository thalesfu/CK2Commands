package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1971] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1971][71971] = People_71971
	HistoryPeopleMap[1971][161971] = People_161971
	HistoryPeopleMap[1971][191971] = People_191971
}
