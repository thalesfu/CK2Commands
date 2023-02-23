package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利多LidoBarony struct {
	feud.BaseBarony
}

var BLido利多 feud.Barony = &利多LidoBarony{}

func init() {
	f := BLido利多.(*利多LidoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lido",
		TitleName: "利多",
		TitleCode: "b_lido",
	}
}
