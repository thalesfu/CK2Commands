package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉罗达LarodaBarony struct {
	feud.BaseBarony
}

var BLaroda拉罗达 feud.Barony = &拉罗达LarodaBarony{}

func init() {
    f := BLaroda拉罗达.(*拉罗达LarodaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laroda",
		TitleName: "拉罗达",
		TitleCode: "b_laroda",
	}
}
