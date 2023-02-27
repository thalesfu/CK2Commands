package kalinganagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利迦俱蓝SrikakulamBarony struct {
	feud.BaseBarony
}

var BSrikakulam室利迦俱蓝 feud.Barony = &室利迦俱蓝SrikakulamBarony{}

func init() {
    f := BSrikakulam室利迦俱蓝.(*室利迦俱蓝SrikakulamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srikakulam",
		TitleName: "室利迦俱蓝",
		TitleCode: "b_srikakulam",
	}
}
