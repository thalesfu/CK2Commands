package coruna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费罗尔FerrolBarony struct {
	feud.BaseBarony
}

var BFerrol费罗尔 feud.Barony = &费罗尔FerrolBarony{}

func init() {
    f := BFerrol费罗尔.(*费罗尔FerrolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ferrol",
		TitleName: "费罗尔",
		TitleCode: "b_ferrol",
	}
}
