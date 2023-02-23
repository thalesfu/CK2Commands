package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯克沃尔KirkwallBarony struct {
	feud.BaseBarony
}

var BKirkwall柯克沃尔 feud.Barony = &柯克沃尔KirkwallBarony{}

func init() {
	f := BKirkwall柯克沃尔.(*柯克沃尔KirkwallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirkwall",
		TitleName: "柯克沃尔",
		TitleCode: "b_kirkwall",
	}
}
