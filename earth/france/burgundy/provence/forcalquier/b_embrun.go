package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂布兰EmbrunBarony struct {
	feud.BaseBarony
}

var BEmbrun昂布兰 feud.Barony = &昂布兰EmbrunBarony{}

func init() {
    f := BEmbrun昂布兰.(*昂布兰EmbrunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "embrun",
		TitleName: "昂布兰",
		TitleCode: "b_embrun",
	}
}
