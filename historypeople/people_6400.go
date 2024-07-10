package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6400][6400] = People_6400
	HistoryPeopleMap[6400][166400] = People_166400
	HistoryPeopleMap[6400][476400] = People_476400
}
