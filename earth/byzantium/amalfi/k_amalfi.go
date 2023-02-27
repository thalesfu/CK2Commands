package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmalfiKingdom interface {
    feud.Kingdom
}

type 阿马尔菲AmalfiKingdom struct {
	feud.BaseKingdom
}

var KAmalfi阿马尔菲 AmalfiKingdom = &阿马尔菲AmalfiKingdom{}

func init() {
	f := KAmalfi阿马尔菲.(*阿马尔菲AmalfiKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "amalfi",
		TitleName: "阿马尔菲",
		TitleCode: "k_amalfi",
		Dukes:  map[string]feud.Duke{},
	}

}
