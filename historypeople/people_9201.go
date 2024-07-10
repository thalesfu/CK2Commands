package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9201] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9201][159201] = People_159201
	HistoryPeopleMap[9201][189201] = People_189201
}
