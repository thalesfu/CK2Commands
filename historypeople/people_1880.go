package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1880] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1880][31880] = People_31880
	HistoryPeopleMap[1880][71880] = People_71880
	HistoryPeopleMap[1880][191880] = People_191880
}
