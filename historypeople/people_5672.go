package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5672] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5672][145672] = People_145672
	HistoryPeopleMap[5672][205672] = People_205672
	HistoryPeopleMap[5672][215672] = People_215672
}
