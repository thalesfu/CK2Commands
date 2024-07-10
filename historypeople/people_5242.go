package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5242] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5242][125242] = People_125242
	HistoryPeopleMap[5242][145242] = People_145242
	HistoryPeopleMap[5242][155242] = People_155242
	HistoryPeopleMap[5242][205242] = People_205242
}
