package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9200][159200] = People_159200
	HistoryPeopleMap[9200][189200] = People_189200
	HistoryPeopleMap[9200][479200] = People_479200
}
