package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维亚纳堡VianadocasteloBarony struct {
	feud.BaseBarony
}

var BVianadocastelo维亚纳堡 feud.Barony = &维亚纳堡VianadocasteloBarony{}

func init() {
	f := BVianadocastelo维亚纳堡.(*维亚纳堡VianadocasteloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vianadocastelo",
		TitleName: "维亚纳堡",
		TitleCode: "b_vianadocastelo",
	}
}
