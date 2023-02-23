package pannagallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 军陀KundaBarony struct {
	feud.BaseBarony
}

var BKunda军陀 feud.Barony = &军陀KundaBarony{}

func init() {
	f := BKunda军陀.(*军陀KundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunda",
		TitleName: "军陀",
		TitleCode: "b_kunda",
	}
}
