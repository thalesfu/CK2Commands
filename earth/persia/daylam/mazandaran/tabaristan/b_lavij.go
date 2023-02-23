package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉瓦基LavijBarony struct {
	feud.BaseBarony
}

var BLavij拉瓦基 feud.Barony = &拉瓦基LavijBarony{}

func init() {
	f := BLavij拉瓦基.(*拉瓦基LavijBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lavij",
		TitleName: "拉瓦基",
		TitleCode: "b_lavij",
	}
}
