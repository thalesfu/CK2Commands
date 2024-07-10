package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1908] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1908][71908] = People_71908
	HistoryPeopleMap[1908][191908] = People_191908
	HistoryPeopleMap[1908][461908] = People_461908
}
