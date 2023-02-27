package austria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AustriaKingdom interface {
    feud.Kingdom
}

type 奥地利AustriaKingdom struct {
	feud.BaseKingdom
}

var KAustria奥地利 AustriaKingdom = &奥地利AustriaKingdom{}

func init() {
	f := KAustria奥地利.(*奥地利AustriaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "austria",
		TitleName: "奥地利",
		TitleCode: "k_austria",
		Dukes:  map[string]feud.Duke{},
	}

}
