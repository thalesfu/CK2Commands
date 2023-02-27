package saruhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaruhanKingdom interface {
    feud.Kingdom
}

type 萨鲁汗SaruhanKingdom struct {
	feud.BaseKingdom
}

var KSaruhan萨鲁汗 SaruhanKingdom = &萨鲁汗SaruhanKingdom{}

func init() {
	f := KSaruhan萨鲁汗.(*萨鲁汗SaruhanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "saruhan",
		TitleName: "萨鲁汗",
		TitleCode: "k_saruhan",
		Dukes:  map[string]feud.Duke{},
	}

}
