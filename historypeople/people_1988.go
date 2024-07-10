package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1988] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1988][71988] = People_71988
	HistoryPeopleMap[1988][191988] = People_191988
}
