package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5599] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5599][145599] = People_145599
	HistoryPeopleMap[5599][205599] = People_205599
	HistoryPeopleMap[5599][215599] = People_215599
}
