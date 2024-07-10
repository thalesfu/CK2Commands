package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1918] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1918][161918] = People_161918
	HistoryPeopleMap[1918][191918] = People_191918
	HistoryPeopleMap[1918][461918] = People_461918
}
