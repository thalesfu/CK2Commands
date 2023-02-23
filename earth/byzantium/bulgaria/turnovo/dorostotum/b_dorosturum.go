package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜罗斯托鲁姆DorosturumBarony struct {
	feud.BaseBarony
}

var BDorosturum杜罗斯托鲁姆 feud.Barony = &杜罗斯托鲁姆DorosturumBarony{}

func init() {
	f := BDorosturum杜罗斯托鲁姆.(*杜罗斯托鲁姆DorosturumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorosturum",
		TitleName: "杜罗斯托鲁姆",
		TitleCode: "b_dorosturum",
	}
}
