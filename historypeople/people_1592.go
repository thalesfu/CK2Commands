package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1592] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1592][41592] = People_41592
	HistoryPeopleMap[1592][71592] = People_71592
	HistoryPeopleMap[1592][191592] = People_191592
}
