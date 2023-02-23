package kharga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰赫达姆GahdamBarony struct {
	feud.BaseBarony
}

var BGahdam杰赫达姆 feud.Barony = &杰赫达姆GahdamBarony{}

func init() {
	f := BGahdam杰赫达姆.(*杰赫达姆GahdamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gahdam",
		TitleName: "杰赫达姆",
		TitleCode: "b_gahdam",
	}
}
