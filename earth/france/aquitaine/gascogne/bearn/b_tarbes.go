package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔布TarbesBarony struct {
	feud.BaseBarony
}

var BTarbes塔布 feud.Barony = &塔布TarbesBarony{}

func init() {
	f := BTarbes塔布.(*塔布TarbesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarbes",
		TitleName: "塔布",
		TitleCode: "b_tarbes",
	}
}
