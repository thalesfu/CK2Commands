package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪韦齐斯DevizesBarony struct {
	feud.BaseBarony
}

var BDevizes迪韦齐斯 feud.Barony = &迪韦齐斯DevizesBarony{}

func init() {
    f := BDevizes迪韦齐斯.(*迪韦齐斯DevizesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devizes",
		TitleName: "迪韦齐斯",
		TitleCode: "b_devizes",
	}
}
