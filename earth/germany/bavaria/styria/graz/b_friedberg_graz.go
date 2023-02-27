package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗里德贝格Friedberg_grazBarony struct {
	feud.BaseBarony
}

var BFriedberg_graz弗里德贝格 feud.Barony = &弗里德贝格Friedberg_grazBarony{}

func init() {
    f := BFriedberg_graz弗里德贝格.(*弗里德贝格Friedberg_grazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "friedberg_graz",
		TitleName: "弗里德贝格",
		TitleCode: "b_friedberg_graz",
	}
}
