package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3888] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3888][73888] = People_73888
}
