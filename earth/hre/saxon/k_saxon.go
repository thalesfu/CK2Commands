package saxon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaxonKingdom interface {
    feud.Kingdom
}

type 萨克森SaxonKingdom struct {
	feud.BaseKingdom
}

var KSaxon萨克森 SaxonKingdom = &萨克森SaxonKingdom{}

func init() {
	f := KSaxon萨克森.(*萨克森SaxonKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "saxon",
		TitleName: "萨克森",
		TitleCode: "k_saxon",
		Dukes:  map[string]feud.Duke{},
	}

}
