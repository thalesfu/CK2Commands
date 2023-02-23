package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅芬NephinBarony struct {
	feud.BaseBarony
}

var BNephin涅芬 feud.Barony = &涅芬NephinBarony{}

func init() {
	f := BNephin涅芬.(*涅芬NephinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nephin",
		TitleName: "涅芬",
		TitleCode: "b_nephin",
	}
}
