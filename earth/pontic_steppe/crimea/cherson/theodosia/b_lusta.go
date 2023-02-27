package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 路斯塔LustaBarony struct {
	feud.BaseBarony
}

var BLusta路斯塔 feud.Barony = &路斯塔LustaBarony{}

func init() {
    f := BLusta路斯塔.(*路斯塔LustaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lusta",
		TitleName: "路斯塔",
		TitleCode: "b_lusta",
	}
}
