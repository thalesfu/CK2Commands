package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普奇塞达PuigcerdaBarony struct {
	feud.BaseBarony
}

var BPuigcerda普奇塞达 feud.Barony = &普奇塞达PuigcerdaBarony{}

func init() {
    f := BPuigcerda普奇塞达.(*普奇塞达PuigcerdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puigcerda",
		TitleName: "普奇塞达",
		TitleCode: "b_puigcerda",
	}
}
