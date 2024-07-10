package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1873] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1873][31873] = People_31873
	HistoryPeopleMap[1873][71873] = People_71873
	HistoryPeopleMap[1873][191873] = People_191873
}
