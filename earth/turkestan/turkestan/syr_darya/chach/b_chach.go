package chach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赭石ChachBarony struct {
	feud.BaseBarony
}

var BChach赭石 feud.Barony = &赭石ChachBarony{}

func init() {
    f := BChach赭石.(*赭石ChachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chach",
		TitleName: "赭石",
		TitleCode: "b_chach",
	}
}
