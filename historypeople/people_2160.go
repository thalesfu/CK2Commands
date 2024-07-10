package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2160] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2160][32160] = People_32160
	HistoryPeopleMap[2160][72160] = People_72160
	HistoryPeopleMap[2160][82160] = People_82160
	HistoryPeopleMap[2160][142160] = People_142160
	HistoryPeopleMap[2160][252160] = People_252160
	HistoryPeopleMap[2160][462160] = People_462160
}
