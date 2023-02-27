package pettau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗里道FriedauBarony struct {
	feud.BaseBarony
}

var BFriedau弗里道 feud.Barony = &弗里道FriedauBarony{}

func init() {
    f := BFriedau弗里道.(*弗里道FriedauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "friedau",
		TitleName: "弗里道",
		TitleCode: "b_friedau",
	}
}
