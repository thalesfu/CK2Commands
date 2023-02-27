package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeonKingdom interface {
    feud.Kingdom
}

type 莱昂LeonKingdom struct {
	feud.BaseKingdom
}

var KLeon莱昂 LeonKingdom = &莱昂LeonKingdom{}

func init() {
	f := KLeon莱昂.(*莱昂LeonKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "leon",
		TitleName: "莱昂",
		TitleCode: "k_leon",
		Dukes:  map[string]feud.Duke{},
	}

}
