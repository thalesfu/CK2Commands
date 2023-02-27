package ottoman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OttomanKingdom interface {
    feud.Kingdom
}

type 奥斯曼OttomanKingdom struct {
	feud.BaseKingdom
}

var KOttoman奥斯曼 OttomanKingdom = &奥斯曼OttomanKingdom{}

func init() {
	f := KOttoman奥斯曼.(*奥斯曼OttomanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "ottoman",
		TitleName: "奥斯曼",
		TitleCode: "k_ottoman",
		Dukes:  map[string]feud.Duke{},
	}

}
