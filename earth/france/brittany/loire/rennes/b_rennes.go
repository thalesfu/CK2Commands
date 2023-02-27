package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷恩RennesBarony struct {
	feud.BaseBarony
}

var BRennes雷恩 feud.Barony = &雷恩RennesBarony{}

func init() {
    f := BRennes雷恩.(*雷恩RennesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rennes",
		TitleName: "雷恩",
		TitleCode: "b_rennes",
	}
}
