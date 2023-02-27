package oleshye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔松DnieprkhersonBarony struct {
	feud.BaseBarony
}

var BDnieprkherson赫尔松 feud.Barony = &赫尔松DnieprkhersonBarony{}

func init() {
    f := BDnieprkherson赫尔松.(*赫尔松DnieprkhersonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dnieprkherson",
		TitleName: "赫尔松",
		TitleCode: "b_dnieprkherson",
	}
}
