package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥里希AurichBarony struct {
	feud.BaseBarony
}

var BAurich奥里希 feud.Barony = &奥里希AurichBarony{}

func init() {
    f := BAurich奥里希.(*奥里希AurichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aurich",
		TitleName: "奥里希",
		TitleCode: "b_aurich",
	}
}
