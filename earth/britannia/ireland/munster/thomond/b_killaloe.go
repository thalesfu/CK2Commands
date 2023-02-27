package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基拉卢KillaloeBarony struct {
	feud.BaseBarony
}

var BKillaloe基拉卢 feud.Barony = &基拉卢KillaloeBarony{}

func init() {
    f := BKillaloe基拉卢.(*基拉卢KillaloeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "killaloe",
		TitleName: "基拉卢",
		TitleCode: "b_killaloe",
	}
}
