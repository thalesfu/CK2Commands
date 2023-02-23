package burhanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伐迦利VaghliBarony struct {
	feud.BaseBarony
}

var BVaghli伐迦利 feud.Barony = &伐迦利VaghliBarony{}

func init() {
	f := BVaghli伐迦利.(*伐迦利VaghliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaghli",
		TitleName: "伐迦利",
		TitleCode: "b_vaghli",
	}
}
