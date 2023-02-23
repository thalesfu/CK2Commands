package vemulavada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补多利PotaliBarony struct {
	feud.BaseBarony
}

var BPotali补多利 feud.Barony = &补多利PotaliBarony{}

func init() {
	f := BPotali补多利.(*补多利PotaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "potali",
		TitleName: "补多利",
		TitleCode: "b_potali",
	}
}
