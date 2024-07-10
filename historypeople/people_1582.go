package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1582] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1582][41582] = People_41582
	HistoryPeopleMap[1582][71582] = People_71582
	HistoryPeopleMap[1582][191582] = People_191582
}
