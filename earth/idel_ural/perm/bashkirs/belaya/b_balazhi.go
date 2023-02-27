package belaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉日BalazhiBarony struct {
	feud.BaseBarony
}

var BBalazhi巴拉日 feud.Barony = &巴拉日BalazhiBarony{}

func init() {
    f := BBalazhi巴拉日.(*巴拉日BalazhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balazhi",
		TitleName: "巴拉日",
		TitleCode: "b_balazhi",
	}
}
