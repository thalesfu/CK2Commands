package otgontenger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴彦布拉格BayanbulagBarony struct {
	feud.BaseBarony
}

var BBayanbulag巴彦布拉格 feud.Barony = &巴彦布拉格BayanbulagBarony{}

func init() {
    f := BBayanbulag巴彦布拉格.(*巴彦布拉格BayanbulagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayanbulag",
		TitleName: "巴彦布拉格",
		TitleCode: "b_bayanbulag",
	}
}
