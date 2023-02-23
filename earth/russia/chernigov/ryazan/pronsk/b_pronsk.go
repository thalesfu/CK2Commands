package pronsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普龙斯克PronskBarony struct {
	feud.BaseBarony
}

var BPronsk普龙斯克 feud.Barony = &普龙斯克PronskBarony{}

func init() {
	f := BPronsk普龙斯克.(*普龙斯克PronskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pronsk",
		TitleName: "普龙斯克",
		TitleCode: "b_pronsk",
	}
}
