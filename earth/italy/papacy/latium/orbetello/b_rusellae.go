package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁塞莱RusellaeBarony struct {
	feud.BaseBarony
}

var BRusellae鲁塞莱 feud.Barony = &鲁塞莱RusellaeBarony{}

func init() {
	f := BRusellae鲁塞莱.(*鲁塞莱RusellaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rusellae",
		TitleName: "鲁塞莱",
		TitleCode: "b_rusellae",
	}
}
