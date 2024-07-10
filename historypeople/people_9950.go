package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9950] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9950][9950] = People_9950
	HistoryPeopleMap[9950][159950] = People_159950
}
