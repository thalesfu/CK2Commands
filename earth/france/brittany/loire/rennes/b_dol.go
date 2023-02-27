package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多勒DolBarony struct {
	feud.BaseBarony
}

var BDol多勒 feud.Barony = &多勒DolBarony{}

func init() {
    f := BDol多勒.(*多勒DolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dol",
		TitleName: "多勒",
		TitleCode: "b_dol",
	}
}
