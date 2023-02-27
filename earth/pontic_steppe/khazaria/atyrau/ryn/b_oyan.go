package ryn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥扬OyanBarony struct {
	feud.BaseBarony
}

var BOyan奥扬 feud.Barony = &奥扬OyanBarony{}

func init() {
    f := BOyan奥扬.(*奥扬OyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oyan",
		TitleName: "奥扬",
		TitleCode: "b_oyan",
	}
}
