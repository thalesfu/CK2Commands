package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5566] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5566][205566] = People_205566
}
