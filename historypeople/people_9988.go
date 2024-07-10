package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9988] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9988][159988] = People_159988
	HistoryPeopleMap[9988][1059988] = People_1059988
}
