package nagormo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达布逊DabuxunBarony struct {
	feud.BaseBarony
}

var BDabuxun达布逊 feud.Barony = &达布逊DabuxunBarony{}

func init() {
    f := BDabuxun达布逊.(*达布逊DabuxunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dabuxun",
		TitleName: "达布逊",
		TitleCode: "b_dabuxun",
	}
}
