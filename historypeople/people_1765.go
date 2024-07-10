package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1765] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1765][1765] = People_1765
	HistoryPeopleMap[1765][31765] = People_31765
	HistoryPeopleMap[1765][71765] = People_71765
	HistoryPeopleMap[1765][161765] = People_161765
	HistoryPeopleMap[1765][191765] = People_191765
}
