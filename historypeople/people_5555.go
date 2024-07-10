package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5555] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5555][205555] = People_205555
	HistoryPeopleMap[5555][215555] = People_215555
	HistoryPeopleMap[5555][455555] = People_455555
	HistoryPeopleMap[5555][465555] = People_465555
	HistoryPeopleMap[5555][475555] = People_475555
}
