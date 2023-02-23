package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旺塔杜尔VentadourBarony struct {
	feud.BaseBarony
}

var BVentadour旺塔杜尔 feud.Barony = &旺塔杜尔VentadourBarony{}

func init() {
	f := BVentadour旺塔杜尔.(*旺塔杜尔VentadourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ventadour",
		TitleName: "旺塔杜尔",
		TitleCode: "b_ventadour",
	}
}
