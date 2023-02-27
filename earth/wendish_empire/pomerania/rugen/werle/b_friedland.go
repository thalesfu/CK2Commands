package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗里德兰FriedlandBarony struct {
	feud.BaseBarony
}

var BFriedland弗里德兰 feud.Barony = &弗里德兰FriedlandBarony{}

func init() {
    f := BFriedland弗里德兰.(*弗里德兰FriedlandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "friedland",
		TitleName: "弗里德兰",
		TitleCode: "b_friedland",
	}
}
