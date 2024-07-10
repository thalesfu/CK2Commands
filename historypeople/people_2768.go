package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2768] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2768][32768] = People_32768
	HistoryPeopleMap[2768][72768] = People_72768
	HistoryPeopleMap[2768][142768] = People_142768
	HistoryPeopleMap[2768][232768] = People_232768
}
