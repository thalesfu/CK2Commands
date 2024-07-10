package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1919] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1919][161919] = People_161919
	HistoryPeopleMap[1919][191919] = People_191919
	HistoryPeopleMap[1919][461919] = People_461919
}
