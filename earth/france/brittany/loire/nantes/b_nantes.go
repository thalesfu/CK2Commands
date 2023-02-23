package nantes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南特NantesBarony struct {
	feud.BaseBarony
}

var BNantes南特 feud.Barony = &南特NantesBarony{}

func init() {
	f := BNantes南特.(*南特NantesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nantes",
		TitleName: "南特",
		TitleCode: "b_nantes",
	}
}
