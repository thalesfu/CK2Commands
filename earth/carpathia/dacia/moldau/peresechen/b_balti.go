package peresechen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯尔兹BaltiBarony struct {
	feud.BaseBarony
}

var BBalti伯尔兹 feud.Barony = &伯尔兹BaltiBarony{}

func init() {
	f := BBalti伯尔兹.(*伯尔兹BaltiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balti",
		TitleName: "伯尔兹",
		TitleCode: "b_balti",
	}
}
