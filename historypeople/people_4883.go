package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4883] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4883][34883] = People_34883
}
