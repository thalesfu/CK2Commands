package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥洛龙OloronBarony struct {
	feud.BaseBarony
}

var BOloron奥洛龙 feud.Barony = &奥洛龙OloronBarony{}

func init() {
    f := BOloron奥洛龙.(*奥洛龙OloronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oloron",
		TitleName: "奥洛龙",
		TitleCode: "b_oloron",
	}
}
