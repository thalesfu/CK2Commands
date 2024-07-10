package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1907] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1907][71907] = People_71907
	HistoryPeopleMap[1907][191907] = People_191907
	HistoryPeopleMap[1907][461907] = People_461907
}
