package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里博尔KrborBarony struct {
	feud.BaseBarony
}

var BKrbor克里博尔 feud.Barony = &克里博尔KrborBarony{}

func init() {
    f := BKrbor克里博尔.(*克里博尔KrborBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krbor",
		TitleName: "克里博尔",
		TitleCode: "b_krbor",
	}
}
