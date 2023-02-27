package gower

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰代洛LlandeiloBarony struct {
	feud.BaseBarony
}

var BLlandeilo兰代洛 feud.Barony = &兰代洛LlandeiloBarony{}

func init() {
    f := BLlandeilo兰代洛.(*兰代洛LlandeiloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llandeilo",
		TitleName: "兰代洛",
		TitleCode: "b_llandeilo",
	}
}
