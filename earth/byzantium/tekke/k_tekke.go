package tekke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TekkeKingdom interface {
    feud.Kingdom
}

type 泰凯TekkeKingdom struct {
	feud.BaseKingdom
}

var KTekke泰凯 TekkeKingdom = &泰凯TekkeKingdom{}

func init() {
	f := KTekke泰凯.(*泰凯TekkeKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "tekke",
		TitleName: "泰凯",
		TitleCode: "k_tekke",
		Dukes:  map[string]feud.Duke{},
	}

}
