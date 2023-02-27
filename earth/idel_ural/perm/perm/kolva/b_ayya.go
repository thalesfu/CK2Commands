package kolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾亚AyyaBarony struct {
	feud.BaseBarony
}

var BAyya艾亚 feud.Barony = &艾亚AyyaBarony{}

func init() {
    f := BAyya艾亚.(*艾亚AyyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayya",
		TitleName: "艾亚",
		TitleCode: "b_ayya",
	}
}
