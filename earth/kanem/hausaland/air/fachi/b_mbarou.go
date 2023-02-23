package fachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆巴鲁MbarouBarony struct {
	feud.BaseBarony
}

var BMbarou姆巴鲁 feud.Barony = &姆巴鲁MbarouBarony{}

func init() {
	f := BMbarou姆巴鲁.(*姆巴鲁MbarouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mbarou",
		TitleName: "姆巴鲁",
		TitleCode: "b_mbarou",
	}
}
