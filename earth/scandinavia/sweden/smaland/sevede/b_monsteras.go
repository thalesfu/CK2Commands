package sevede

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙斯特罗斯MonsterasBarony struct {
	feud.BaseBarony
}

var BMonsteras蒙斯特罗斯 feud.Barony = &蒙斯特罗斯MonsterasBarony{}

func init() {
	f := BMonsteras蒙斯特罗斯.(*蒙斯特罗斯MonsterasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monsteras",
		TitleName: "蒙斯特罗斯",
		TitleCode: "b_monsteras",
	}
}
