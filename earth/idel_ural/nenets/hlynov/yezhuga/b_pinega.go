package yezhuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮涅加PinegaBarony struct {
	feud.BaseBarony
}

var BPinega皮涅加 feud.Barony = &皮涅加PinegaBarony{}

func init() {
    f := BPinega皮涅加.(*皮涅加PinegaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pinega",
		TitleName: "皮涅加",
		TitleCode: "b_pinega",
	}
}
