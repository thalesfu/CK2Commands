package eastern_marches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Eastern_marchesKingdom interface {
    feud.Kingdom
}

type 奥地利Eastern_marchesKingdom struct {
	feud.BaseKingdom
}

var KEastern_marches奥地利 Eastern_marchesKingdom = &奥地利Eastern_marchesKingdom{}

func init() {
	f := KEastern_marches奥地利.(*奥地利Eastern_marchesKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "eastern_marches",
		TitleName: "奥地利",
		TitleCode: "k_eastern_marches",
		Dukes:  map[string]feud.Duke{},
	}

}
