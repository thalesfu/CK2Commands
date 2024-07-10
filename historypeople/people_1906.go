package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1906] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1906][71906] = People_71906
	HistoryPeopleMap[1906][191906] = People_191906
	HistoryPeopleMap[1906][461906] = People_461906
}
