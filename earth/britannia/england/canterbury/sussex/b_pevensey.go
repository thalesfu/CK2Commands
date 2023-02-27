package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩文西PevenseyBarony struct {
	feud.BaseBarony
}

var BPevensey佩文西 feud.Barony = &佩文西PevenseyBarony{}

func init() {
    f := BPevensey佩文西.(*佩文西PevenseyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pevensey",
		TitleName: "佩文西",
		TitleCode: "b_pevensey",
	}
}
