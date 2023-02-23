package kamatapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦摩多城KamatapurBarony struct {
	feud.BaseBarony
}

var BKamatapur迦摩多城 feud.Barony = &迦摩多城KamatapurBarony{}

func init() {
	f := BKamatapur迦摩多城.(*迦摩多城KamatapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamatapur",
		TitleName: "迦摩多城",
		TitleCode: "b_kamatapur",
	}
}
