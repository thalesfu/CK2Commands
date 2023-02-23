package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗里德贝格FriedbergBarony struct {
	feud.BaseBarony
}

var BFriedberg弗里德贝格 feud.Barony = &弗里德贝格FriedbergBarony{}

func init() {
	f := BFriedberg弗里德贝格.(*弗里德贝格FriedbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "friedberg",
		TitleName: "弗里德贝格",
		TitleCode: "b_friedberg",
	}
}
