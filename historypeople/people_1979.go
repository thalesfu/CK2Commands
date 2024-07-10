package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1979] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1979][71979] = People_71979
	HistoryPeopleMap[1979][161979] = People_161979
	HistoryPeopleMap[1979][191979] = People_191979
}
