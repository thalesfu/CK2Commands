package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1977] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1977][71977] = People_71977
	HistoryPeopleMap[1977][161977] = People_161977
	HistoryPeopleMap[1977][191977] = People_191977
}
