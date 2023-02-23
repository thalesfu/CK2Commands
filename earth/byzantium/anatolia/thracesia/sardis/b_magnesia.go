package sardis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格涅西亚MagnesiaBarony struct {
	feud.BaseBarony
}

var BMagnesia马格涅西亚 feud.Barony = &马格涅西亚MagnesiaBarony{}

func init() {
	f := BMagnesia马格涅西亚.(*马格涅西亚MagnesiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "magnesia",
		TitleName: "马格涅西亚",
		TitleCode: "b_magnesia",
	}
}
