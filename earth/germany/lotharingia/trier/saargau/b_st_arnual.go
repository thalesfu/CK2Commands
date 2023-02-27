package saargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣阿努亚尔St_arnualBarony struct {
	feud.BaseBarony
}

var BSt_arnual圣阿努亚尔 feud.Barony = &圣阿努亚尔St_arnualBarony{}

func init() {
    f := BSt_arnual圣阿努亚尔.(*圣阿努亚尔St_arnualBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_arnual",
		TitleName: "圣阿努亚尔",
		TitleCode: "b_st_arnual",
	}
}
