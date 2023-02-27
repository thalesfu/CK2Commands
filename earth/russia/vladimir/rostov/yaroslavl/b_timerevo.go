package yaroslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季梅列沃TimerevoBarony struct {
	feud.BaseBarony
}

var BTimerevo季梅列沃 feud.Barony = &季梅列沃TimerevoBarony{}

func init() {
    f := BTimerevo季梅列沃.(*季梅列沃TimerevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timerevo",
		TitleName: "季梅列沃",
		TitleCode: "b_timerevo",
	}
}
