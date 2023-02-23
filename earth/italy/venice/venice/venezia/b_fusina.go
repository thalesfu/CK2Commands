package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富西纳FusinaBarony struct {
	feud.BaseBarony
}

var BFusina富西纳 feud.Barony = &富西纳FusinaBarony{}

func init() {
	f := BFusina富西纳.(*富西纳FusinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fusina",
		TitleName: "富西纳",
		TitleCode: "b_fusina",
	}
}
