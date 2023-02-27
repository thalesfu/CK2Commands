package romny

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃热加VozhegaBarony struct {
	feud.BaseBarony
}

var BVozhega沃热加 feud.Barony = &沃热加VozhegaBarony{}

func init() {
    f := BVozhega沃热加.(*沃热加VozhegaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vozhega",
		TitleName: "沃热加",
		TitleCode: "b_vozhega",
	}
}
