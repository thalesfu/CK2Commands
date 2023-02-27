package lettigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 高耶纳GaujienaBarony struct {
	feud.BaseBarony
}

var BGaujiena高耶纳 feud.Barony = &高耶纳GaujienaBarony{}

func init() {
    f := BGaujiena高耶纳.(*高耶纳GaujienaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaujiena",
		TitleName: "高耶纳",
		TitleCode: "b_gaujiena",
	}
}
