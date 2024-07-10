package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6832] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6832][6832] = People_6832
	HistoryPeopleMap[6832][166832] = People_166832
}
