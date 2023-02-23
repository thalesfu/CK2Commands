package rajanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀迦DajalBarony struct {
	feud.BaseBarony
}

var BDajal陀迦 feud.Barony = &陀迦DajalBarony{}

func init() {
	f := BDajal陀迦.(*陀迦DajalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dajal",
		TitleName: "陀迦",
		TitleCode: "b_dajal",
	}
}
