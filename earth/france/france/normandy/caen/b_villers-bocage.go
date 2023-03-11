package caen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维莱博卡日Villers_bocageBarony struct {
	feud.BaseBarony
}

var BVillers_bocage维莱博卡日 feud.Barony = &维莱博卡日Villers_bocageBarony{}

func init() {
    f := BVillers_bocage维莱博卡日.(*维莱博卡日Villers_bocageBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villers_bocage",
		TitleName: "维莱博卡日",
		TitleCode: "b_villers-bocage",
	}
}
