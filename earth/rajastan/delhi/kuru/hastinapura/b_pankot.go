package hastinapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般拘吒PankotBarony struct {
	feud.BaseBarony
}

var BPankot般拘吒 feud.Barony = &般拘吒PankotBarony{}

func init() {
	f := BPankot般拘吒.(*般拘吒PankotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pankot",
		TitleName: "般拘吒",
		TitleCode: "b_pankot",
	}
}
