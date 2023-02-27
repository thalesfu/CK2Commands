package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪巴DibbaBarony struct {
	feud.BaseBarony
}

var BDibba迪巴 feud.Barony = &迪巴DibbaBarony{}

func init() {
    f := BDibba迪巴.(*迪巴DibbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dibba",
		TitleName: "迪巴",
		TitleCode: "b_dibba",
	}
}
