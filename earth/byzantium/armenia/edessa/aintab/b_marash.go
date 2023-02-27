package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉什MarashBarony struct {
	feud.BaseBarony
}

var BMarash马拉什 feud.Barony = &马拉什MarashBarony{}

func init() {
    f := BMarash马拉什.(*马拉什MarashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marash",
		TitleName: "马拉什",
		TitleCode: "b_marash",
	}
}
