package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5623] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5623][145623] = People_145623
	HistoryPeopleMap[5623][205623] = People_205623
}
