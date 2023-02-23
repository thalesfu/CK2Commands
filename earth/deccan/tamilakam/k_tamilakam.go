package tamilakam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TamilakamKingdom interface {
	feud.Kingdom
}

type 昙密罗迦摩TamilakamKingdom struct {
	feud.BaseKingdom
}

var KTamilakam昙密罗迦摩 TamilakamKingdom = &昙密罗迦摩TamilakamKingdom{}

func init() {
	f := KTamilakam昙密罗迦摩.(*昙密罗迦摩TamilakamKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "tamilakam",
		TitleName: "昙密罗迦摩",
		TitleCode: "k_tamilakam",
		Dukes:     map[string]feud.Duke{},
	}

}
