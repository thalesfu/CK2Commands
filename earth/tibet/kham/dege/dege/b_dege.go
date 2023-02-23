package dege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德格DegeBarony struct {
	feud.BaseBarony
}

var BDege德格 feud.Barony = &德格DegeBarony{}

func init() {
	f := BDege德格.(*德格DegeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dege",
		TitleName: "德格",
		TitleCode: "b_dege",
	}
}
