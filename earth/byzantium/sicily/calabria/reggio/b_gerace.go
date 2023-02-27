package reggio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰拉切GeraceBarony struct {
	feud.BaseBarony
}

var BGerace杰拉切 feud.Barony = &杰拉切GeraceBarony{}

func init() {
    f := BGerace杰拉切.(*杰拉切GeraceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gerace",
		TitleName: "杰拉切",
		TitleCode: "b_gerace",
	}
}
