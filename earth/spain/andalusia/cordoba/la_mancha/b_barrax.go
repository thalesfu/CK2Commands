package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉克斯BarraxBarony struct {
	feud.BaseBarony
}

var BBarrax巴拉克斯 feud.Barony = &巴拉克斯BarraxBarony{}

func init() {
    f := BBarrax巴拉克斯.(*巴拉克斯BarraxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barrax",
		TitleName: "巴拉克斯",
		TitleCode: "b_barrax",
	}
}
