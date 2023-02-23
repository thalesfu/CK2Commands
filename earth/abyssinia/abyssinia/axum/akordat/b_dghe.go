package akordat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德赫DgheBarony struct {
	feud.BaseBarony
}

var BDghe德赫 feud.Barony = &德赫DgheBarony{}

func init() {
	f := BDghe德赫.(*德赫DgheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dghe",
		TitleName: "德赫",
		TitleCode: "b_dghe",
	}
}
