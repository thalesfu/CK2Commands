package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗罗卡VrakaBarony struct {
	feud.BaseBarony
}

var BVraka弗罗卡 feud.Barony = &弗罗卡VrakaBarony{}

func init() {
    f := BVraka弗罗卡.(*弗罗卡VrakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vraka",
		TitleName: "弗罗卡",
		TitleCode: "b_vraka",
	}
}
