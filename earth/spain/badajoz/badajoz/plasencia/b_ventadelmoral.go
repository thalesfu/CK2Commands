package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本塔德尔莫拉尔VentadelmoralBarony struct {
	feud.BaseBarony
}

var BVentadelmoral本塔德尔莫拉尔 feud.Barony = &本塔德尔莫拉尔VentadelmoralBarony{}

func init() {
	f := BVentadelmoral本塔德尔莫拉尔.(*本塔德尔莫拉尔VentadelmoralBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ventadelmoral",
		TitleName: "本塔德尔莫拉尔",
		TitleCode: "b_ventadelmoral",
	}
}
