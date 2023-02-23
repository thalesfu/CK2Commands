package dwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德温DvinBarony struct {
	feud.BaseBarony
}

var BDvin德温 feud.Barony = &德温DvinBarony{}

func init() {
	f := BDvin德温.(*德温DvinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dvin",
		TitleName: "德温",
		TitleCode: "b_dvin",
	}
}
