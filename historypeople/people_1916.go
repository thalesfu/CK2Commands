package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1916] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1916][161916] = People_161916
	HistoryPeopleMap[1916][191916] = People_191916
	HistoryPeopleMap[1916][461916] = People_461916
}
