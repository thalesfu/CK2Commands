package vorotynsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃罗滕斯克VorotynskBarony struct {
	feud.BaseBarony
}

var BVorotynsk沃罗滕斯克 feud.Barony = &沃罗滕斯克VorotynskBarony{}

func init() {
	f := BVorotynsk沃罗滕斯克.(*沃罗滕斯克VorotynskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vorotynsk",
		TitleName: "沃罗滕斯克",
		TitleCode: "b_vorotynsk",
	}
}
