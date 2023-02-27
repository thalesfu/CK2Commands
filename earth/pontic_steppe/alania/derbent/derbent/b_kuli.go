package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库利KuliBarony struct {
	feud.BaseBarony
}

var BKuli库利 feud.Barony = &库利KuliBarony{}

func init() {
    f := BKuli库利.(*库利KuliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuli",
		TitleName: "库利",
		TitleCode: "b_kuli",
	}
}
