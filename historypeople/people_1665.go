package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1665] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1665][31665] = People_31665
	HistoryPeopleMap[1665][71665] = People_71665
	HistoryPeopleMap[1665][161665] = People_161665
	HistoryPeopleMap[1665][191665] = People_191665
}
