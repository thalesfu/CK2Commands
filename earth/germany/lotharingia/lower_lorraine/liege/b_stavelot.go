package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔沃洛StavelotBarony struct {
	feud.BaseBarony
}

var BStavelot斯塔沃洛 feud.Barony = &斯塔沃洛StavelotBarony{}

func init() {
    f := BStavelot斯塔沃洛.(*斯塔沃洛StavelotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stavelot",
		TitleName: "斯塔沃洛",
		TitleCode: "b_stavelot",
	}
}
