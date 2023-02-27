package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达罗DurrowBarony struct {
	feud.BaseBarony
}

var BDurrow达罗 feud.Barony = &达罗DurrowBarony{}

func init() {
    f := BDurrow达罗.(*达罗DurrowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durrow",
		TitleName: "达罗",
		TitleCode: "b_durrow",
	}
}
