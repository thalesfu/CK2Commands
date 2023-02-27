package hama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉BaraBarony struct {
	feud.BaseBarony
}

var BBara巴拉 feud.Barony = &巴拉BaraBarony{}

func init() {
    f := BBara巴拉.(*巴拉BaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bara",
		TitleName: "巴拉",
		TitleCode: "b_bara",
	}
}
