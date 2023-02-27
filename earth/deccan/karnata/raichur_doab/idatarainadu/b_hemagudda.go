package idatarainadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 醯摩矩吒HemaguddaBarony struct {
	feud.BaseBarony
}

var BHemagudda醯摩矩吒 feud.Barony = &醯摩矩吒HemaguddaBarony{}

func init() {
    f := BHemagudda醯摩矩吒.(*醯摩矩吒HemaguddaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hemagudda",
		TitleName: "醯摩矩吒",
		TitleCode: "b_hemagudda",
	}
}
