package overijssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代芬特尔DeventerBarony struct {
	feud.BaseBarony
}

var BDeventer代芬特尔 feud.Barony = &代芬特尔DeventerBarony{}

func init() {
    f := BDeventer代芬特尔.(*代芬特尔DeventerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deventer",
		TitleName: "代芬特尔",
		TitleCode: "b_deventer",
	}
}
