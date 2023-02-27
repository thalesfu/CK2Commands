package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲雅克FigeacBarony struct {
	feud.BaseBarony
}

var BFigeac菲雅克 feud.Barony = &菲雅克FigeacBarony{}

func init() {
    f := BFigeac菲雅克.(*菲雅克FigeacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "figeac",
		TitleName: "菲雅克",
		TitleCode: "b_figeac",
	}
}
