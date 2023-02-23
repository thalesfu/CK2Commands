package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克朗斯ClonesBarony struct {
	feud.BaseBarony
}

var BClones克朗斯 feud.Barony = &克朗斯ClonesBarony{}

func init() {
	f := BClones克朗斯.(*克朗斯ClonesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clones",
		TitleName: "克朗斯",
		TitleCode: "b_clones",
	}
}
