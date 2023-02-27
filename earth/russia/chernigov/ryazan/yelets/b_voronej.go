package yelets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃罗涅日VoronejBarony struct {
	feud.BaseBarony
}

var BVoronej沃罗涅日 feud.Barony = &沃罗涅日VoronejBarony{}

func init() {
    f := BVoronej沃罗涅日.(*沃罗涅日VoronejBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voronej",
		TitleName: "沃罗涅日",
		TitleCode: "b_voronej",
	}
}
