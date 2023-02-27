package taurica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TauricaKingdom interface {
    feud.Kingdom
}

type 陶里加TauricaKingdom struct {
	feud.BaseKingdom
}

var KTaurica陶里加 TauricaKingdom = &陶里加TauricaKingdom{}

func init() {
	f := KTaurica陶里加.(*陶里加TauricaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "taurica",
		TitleName: "陶里加",
		TitleCode: "k_taurica",
		Dukes:  map[string]feud.Duke{},
	}

}
