package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉亨尼PrachenBarony struct {
	feud.BaseBarony
}

var BPrachen普拉亨尼 feud.Barony = &普拉亨尼PrachenBarony{}

func init() {
    f := BPrachen普拉亨尼.(*普拉亨尼PrachenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prachen",
		TitleName: "普拉亨尼",
		TitleCode: "b_prachen",
	}
}
