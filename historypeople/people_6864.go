package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6864] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6864][6864] = People_6864
	HistoryPeopleMap[6864][166864] = People_166864
}
