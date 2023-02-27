package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲施豪森FischhausenBarony struct {
	feud.BaseBarony
}

var BFischhausen菲施豪森 feud.Barony = &菲施豪森FischhausenBarony{}

func init() {
    f := BFischhausen菲施豪森.(*菲施豪森FischhausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fischhausen",
		TitleName: "菲施豪森",
		TitleCode: "b_fischhausen",
	}
}
