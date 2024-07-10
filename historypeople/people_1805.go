package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1805] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1805][31805] = People_31805
	HistoryPeopleMap[1805][71805] = People_71805
	HistoryPeopleMap[1805][191805] = People_191805
}
