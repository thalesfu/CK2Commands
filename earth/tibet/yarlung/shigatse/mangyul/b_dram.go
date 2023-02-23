package mangyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 樟木DramBarony struct {
	feud.BaseBarony
}

var BDram樟木 feud.Barony = &樟木DramBarony{}

func init() {
	f := BDram樟木.(*樟木DramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dram",
		TitleName: "樟木",
		TitleCode: "b_dram",
	}
}
