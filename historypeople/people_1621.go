package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1621] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1621][1621] = People_1621
	HistoryPeopleMap[1621][71621] = People_71621
	HistoryPeopleMap[1621][161621] = People_161621
	HistoryPeopleMap[1621][191621] = People_191621
}
