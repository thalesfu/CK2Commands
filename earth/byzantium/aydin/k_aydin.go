package aydin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AydinKingdom interface {
	feud.Kingdom
}

type 艾登AydinKingdom struct {
	feud.BaseKingdom
}

var KAydin艾登 AydinKingdom = &艾登AydinKingdom{}

func init() {
	f := KAydin艾登.(*艾登AydinKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "aydin",
		TitleName: "艾登",
		TitleCode: "k_aydin",
		Dukes:     map[string]feud.Duke{},
	}

}
