package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯特明讷KertemindeBarony struct {
	feud.BaseBarony
}

var BKerteminde凯特明讷 feud.Barony = &凯特明讷KertemindeBarony{}

func init() {
    f := BKerteminde凯特明讷.(*凯特明讷KertemindeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerteminde",
		TitleName: "凯特明讷",
		TitleCode: "b_kerteminde",
	}
}
