package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿泽吉阿亚AzeeziayaBarony struct {
	feud.BaseBarony
}

var BAzeeziaya阿泽吉阿亚 feud.Barony = &阿泽吉阿亚AzeeziayaBarony{}

func init() {
    f := BAzeeziaya阿泽吉阿亚.(*阿泽吉阿亚AzeeziayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azeeziaya",
		TitleName: "阿泽吉阿亚",
		TitleCode: "b_azeeziaya",
	}
}
