package ryn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷恩RynBarony struct {
	feud.BaseBarony
}

var BRyn雷恩 feud.Barony = &雷恩RynBarony{}

func init() {
    f := BRyn雷恩.(*雷恩RynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ryn",
		TitleName: "雷恩",
		TitleCode: "b_ryn",
	}
}
