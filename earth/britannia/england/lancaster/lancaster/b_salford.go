package lancaster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔福德SalfordBarony struct {
	feud.BaseBarony
}

var BSalford索尔福德 feud.Barony = &索尔福德SalfordBarony{}

func init() {
	f := BSalford索尔福德.(*索尔福德SalfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salford",
		TitleName: "索尔福德",
		TitleCode: "b_salford",
	}
}
