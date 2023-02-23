package nice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂蒂布AntibesBarony struct {
	feud.BaseBarony
}

var BAntibes昂蒂布 feud.Barony = &昂蒂布AntibesBarony{}

func init() {
	f := BAntibes昂蒂布.(*昂蒂布AntibesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antibes",
		TitleName: "昂蒂布",
		TitleCode: "b_antibes",
	}
}
