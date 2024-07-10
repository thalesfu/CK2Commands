package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5245] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5245][125245] = People_125245
	HistoryPeopleMap[5245][145245] = People_145245
	HistoryPeopleMap[5245][155245] = People_155245
	HistoryPeopleMap[5245][205245] = People_205245
}
