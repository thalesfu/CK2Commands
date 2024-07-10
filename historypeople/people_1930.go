package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1930] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1930][71930] = People_71930
	HistoryPeopleMap[1930][161930] = People_161930
	HistoryPeopleMap[1930][191930] = People_191930
}
