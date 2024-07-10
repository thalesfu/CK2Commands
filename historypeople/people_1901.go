package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1901] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1901][71901] = People_71901
	HistoryPeopleMap[1901][161901] = People_161901
	HistoryPeopleMap[1901][191901] = People_191901
	HistoryPeopleMap[1901][461901] = People_461901
}
