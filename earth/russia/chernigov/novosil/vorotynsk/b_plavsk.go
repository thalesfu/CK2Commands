package vorotynsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉夫斯克PlavskBarony struct {
	feud.BaseBarony
}

var BPlavsk普拉夫斯克 feud.Barony = &普拉夫斯克PlavskBarony{}

func init() {
	f := BPlavsk普拉夫斯克.(*普拉夫斯克PlavskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plavsk",
		TitleName: "普拉夫斯克",
		TitleCode: "b_plavsk",
	}
}
