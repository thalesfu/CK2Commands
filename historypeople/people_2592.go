package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2592] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2592][32592] = People_32592
	HistoryPeopleMap[2592][72592] = People_72592
	HistoryPeopleMap[2592][142592] = People_142592
	HistoryPeopleMap[2592][212592] = People_212592
	HistoryPeopleMap[2592][222592] = People_222592
	HistoryPeopleMap[2592][232592] = People_232592
}
