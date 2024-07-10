package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4777] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4777][74777] = People_74777
}
