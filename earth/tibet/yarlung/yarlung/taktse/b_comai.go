package taktse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 措美ComaiBarony struct {
	feud.BaseBarony
}

var BComai措美 feud.Barony = &措美ComaiBarony{}

func init() {
    f := BComai措美.(*措美ComaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "comai",
		TitleName: "措美",
		TitleCode: "b_comai",
	}
}
