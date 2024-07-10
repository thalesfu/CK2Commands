package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1876] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1876][71876] = People_71876
	HistoryPeopleMap[1876][191876] = People_191876
}
