package cebta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥莱法AulefBarony struct {
	feud.BaseBarony
}

var BAulef奥莱法 feud.Barony = &奥莱法AulefBarony{}

func init() {
	f := BAulef奥莱法.(*奥莱法AulefBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aulef",
		TitleName: "奥莱法",
		TitleCode: "b_aulef",
	}
}
