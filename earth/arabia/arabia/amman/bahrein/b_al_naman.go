package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳曼Al_namanBarony struct {
	feud.BaseBarony
}

var BAl_naman纳曼 feud.Barony = &纳曼Al_namanBarony{}

func init() {
    f := BAl_naman纳曼.(*纳曼Al_namanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_naman",
		TitleName: "纳曼",
		TitleCode: "b_al_naman",
	}
}
