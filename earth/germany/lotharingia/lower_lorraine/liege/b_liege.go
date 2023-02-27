package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列日LiegeBarony struct {
	feud.BaseBarony
}

var BLiege列日 feud.Barony = &列日LiegeBarony{}

func init() {
    f := BLiege列日.(*列日LiegeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liege",
		TitleName: "列日",
		TitleCode: "b_liege",
	}
}
