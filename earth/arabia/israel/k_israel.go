package israel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IsraelKingdom interface {
    feud.Kingdom
}

type 以色列IsraelKingdom struct {
	feud.BaseKingdom
}

var KIsrael以色列 IsraelKingdom = &以色列IsraelKingdom{}

func init() {
	f := KIsrael以色列.(*以色列IsraelKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "israel",
		TitleName: "以色列",
		TitleCode: "k_israel",
		Dukes:  map[string]feud.Duke{},
	}

}
