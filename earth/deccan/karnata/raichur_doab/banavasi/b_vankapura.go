package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 房迦补罗VankapuraBarony struct {
	feud.BaseBarony
}

var BVankapura房迦补罗 feud.Barony = &房迦补罗VankapuraBarony{}

func init() {
    f := BVankapura房迦补罗.(*房迦补罗VankapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vankapura",
		TitleName: "房迦补罗",
		TitleCode: "b_vankapura",
	}
}
