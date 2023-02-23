package forde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗勒FordeBarony struct {
	feud.BaseBarony
}

var BForde弗勒 feud.Barony = &弗勒FordeBarony{}

func init() {
	f := BForde弗勒.(*弗勒FordeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forde",
		TitleName: "弗勒",
		TitleCode: "b_forde",
	}
}
