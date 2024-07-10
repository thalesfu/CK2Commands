package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5234] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5234][125234] = People_125234
	HistoryPeopleMap[5234][145234] = People_145234
	HistoryPeopleMap[5234][155234] = People_155234
	HistoryPeopleMap[5234][205234] = People_205234
}
