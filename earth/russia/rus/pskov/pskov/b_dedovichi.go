package pskov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰多维奇DedovichiBarony struct {
	feud.BaseBarony
}

var BDedovichi杰多维奇 feud.Barony = &杰多维奇DedovichiBarony{}

func init() {
    f := BDedovichi杰多维奇.(*杰多维奇DedovichiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dedovichi",
		TitleName: "杰多维奇",
		TitleCode: "b_dedovichi",
	}
}
