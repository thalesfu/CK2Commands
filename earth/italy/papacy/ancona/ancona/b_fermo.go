package ancona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费尔莫FermoBarony struct {
	feud.BaseBarony
}

var BFermo费尔莫 feud.Barony = &费尔莫FermoBarony{}

func init() {
    f := BFermo费尔莫.(*费尔莫FermoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fermo",
		TitleName: "费尔莫",
		TitleCode: "b_fermo",
	}
}
