package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯基尔斯图纳EskilstunaBarony struct {
	feud.BaseBarony
}

var BEskilstuna埃斯基尔斯图纳 feud.Barony = &埃斯基尔斯图纳EskilstunaBarony{}

func init() {
    f := BEskilstuna埃斯基尔斯图纳.(*埃斯基尔斯图纳EskilstunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eskilstuna",
		TitleName: "埃斯基尔斯图纳",
		TitleCode: "b_eskilstuna",
	}
}
