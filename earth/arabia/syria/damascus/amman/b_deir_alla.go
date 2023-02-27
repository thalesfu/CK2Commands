package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代尔安拉Deir_allaBarony struct {
	feud.BaseBarony
}

var BDeir_alla代尔安拉 feud.Barony = &代尔安拉Deir_allaBarony{}

func init() {
    f := BDeir_alla代尔安拉.(*代尔安拉Deir_allaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deir_alla",
		TitleName: "代尔安拉",
		TitleCode: "b_deir_alla",
	}
}
