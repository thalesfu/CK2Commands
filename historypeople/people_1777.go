package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1777] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1777][31777] = People_31777
	HistoryPeopleMap[1777][161777] = People_161777
	HistoryPeopleMap[1777][191777] = People_191777
}
