package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1889] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1889][31889] = People_31889
	HistoryPeopleMap[1889][71889] = People_71889
	HistoryPeopleMap[1889][191889] = People_191889
}
