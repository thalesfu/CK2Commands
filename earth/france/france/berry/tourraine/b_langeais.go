package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗热LangeaisBarony struct {
	feud.BaseBarony
}

var BLangeais朗热 feud.Barony = &朗热LangeaisBarony{}

func init() {
    f := BLangeais朗热.(*朗热LangeaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "langeais",
		TitleName: "朗热",
		TitleCode: "b_langeais",
	}
}
