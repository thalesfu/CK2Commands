package kubera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 葛仙庙GexianmiaoBarony struct {
	feud.BaseBarony
}

var BGexianmiao葛仙庙 feud.Barony = &葛仙庙GexianmiaoBarony{}

func init() {
	f := BGexianmiao葛仙庙.(*葛仙庙GexianmiaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gexianmiao",
		TitleName: "葛仙庙",
		TitleCode: "b_gexianmiao",
	}
}
