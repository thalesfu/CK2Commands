package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦利格勒VeligradBarony struct {
	feud.BaseBarony
}

var BVeligrad韦利格勒 feud.Barony = &韦利格勒VeligradBarony{}

func init() {
    f := BVeligrad韦利格勒.(*韦利格勒VeligradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veligrad",
		TitleName: "韦利格勒",
		TitleCode: "b_veligrad",
	}
}
