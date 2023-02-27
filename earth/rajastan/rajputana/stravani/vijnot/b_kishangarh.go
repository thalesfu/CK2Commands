package vijnot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 计霜姞利呬KishangarhBarony struct {
	feud.BaseBarony
}

var BKishangarh计霜姞利呬 feud.Barony = &计霜姞利呬KishangarhBarony{}

func init() {
    f := BKishangarh计霜姞利呬.(*计霜姞利呬KishangarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kishangarh",
		TitleName: "计霜姞利呬",
		TitleCode: "b_kishangarh",
	}
}
