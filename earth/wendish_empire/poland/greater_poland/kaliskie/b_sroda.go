package kaliskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希罗达SrodaBarony struct {
	feud.BaseBarony
}

var BSroda希罗达 feud.Barony = &希罗达SrodaBarony{}

func init() {
    f := BSroda希罗达.(*希罗达SrodaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sroda",
		TitleName: "希罗达",
		TitleCode: "b_sroda",
	}
}
