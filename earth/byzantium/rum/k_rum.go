package rum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RumKingdom interface {
	feud.Kingdom
}

type 罗姆RumKingdom struct {
	feud.BaseKingdom
}

var KRum罗姆 RumKingdom = &罗姆RumKingdom{}

func init() {
	f := KRum罗姆.(*罗姆RumKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "rum",
		TitleName: "罗姆",
		TitleCode: "k_rum",
		Dukes:     map[string]feud.Duke{},
	}

}
