package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜勒加Al_balqaBarony struct {
	feud.BaseBarony
}

var BAl_balqa拜勒加 feud.Barony = &拜勒加Al_balqaBarony{}

func init() {
    f := BAl_balqa拜勒加.(*拜勒加Al_balqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_balqa",
		TitleName: "拜勒加",
		TitleCode: "b_al_balqa",
	}
}
