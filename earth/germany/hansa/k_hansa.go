package hansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HansaKingdom interface {
    feud.Kingdom
}

type 汉萨同盟HansaKingdom struct {
	feud.BaseKingdom
}

var KHansa汉萨同盟 HansaKingdom = &汉萨同盟HansaKingdom{}

func init() {
	f := KHansa汉萨同盟.(*汉萨同盟HansaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "hansa",
		TitleName: "汉萨同盟",
		TitleCode: "k_hansa",
		Dukes:  map[string]feud.Duke{},
	}

}
