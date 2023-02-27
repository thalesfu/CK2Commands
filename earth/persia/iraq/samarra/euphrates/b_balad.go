package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜莱德BaladBarony struct {
	feud.BaseBarony
}

var BBalad拜莱德 feud.Barony = &拜莱德BaladBarony{}

func init() {
    f := BBalad拜莱德.(*拜莱德BaladBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balad",
		TitleName: "拜莱德",
		TitleCode: "b_balad",
	}
}
