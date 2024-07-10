package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6845] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6845][6845] = People_6845
	HistoryPeopleMap[6845][166845] = People_166845
	HistoryPeopleMap[6845][206845] = People_206845
}
