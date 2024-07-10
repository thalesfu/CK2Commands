package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6865] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6865][6865] = People_6865
	HistoryPeopleMap[6865][166865] = People_166865
}
