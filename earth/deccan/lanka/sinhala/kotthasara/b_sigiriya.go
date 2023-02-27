package kotthasara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯耆利耶SigiriyaBarony struct {
	feud.BaseBarony
}

var BSigiriya斯耆利耶 feud.Barony = &斯耆利耶SigiriyaBarony{}

func init() {
    f := BSigiriya斯耆利耶.(*斯耆利耶SigiriyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sigiriya",
		TitleName: "斯耆利耶",
		TitleCode: "b_sigiriya",
	}
}
