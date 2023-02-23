package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙斯MonsBarony struct {
	feud.BaseBarony
}

var BMons蒙斯 feud.Barony = &蒙斯MonsBarony{}

func init() {
	f := BMons蒙斯.(*蒙斯MonsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mons",
		TitleName: "蒙斯",
		TitleCode: "b_mons",
	}
}
