package naimisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补婆谛DurbateBarony struct {
	feud.BaseBarony
}

var BDurbate补婆谛 feud.Barony = &补婆谛DurbateBarony{}

func init() {
    f := BDurbate补婆谛.(*补婆谛DurbateBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durbate",
		TitleName: "补婆谛",
		TitleCode: "b_durbate",
	}
}
