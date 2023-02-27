package almada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔梅拉PalmelaBarony struct {
	feud.BaseBarony
}

var BPalmela帕尔梅拉 feud.Barony = &帕尔梅拉PalmelaBarony{}

func init() {
    f := BPalmela帕尔梅拉.(*帕尔梅拉PalmelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palmela",
		TitleName: "帕尔梅拉",
		TitleCode: "b_palmela",
	}
}
