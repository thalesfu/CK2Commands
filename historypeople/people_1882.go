package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1882] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1882][31882] = People_31882
	HistoryPeopleMap[1882][71882] = People_71882
	HistoryPeopleMap[1882][191882] = People_191882
}
