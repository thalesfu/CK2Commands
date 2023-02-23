package orsha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什克洛夫ShklowBarony struct {
	feud.BaseBarony
}

var BShklow什克洛夫 feud.Barony = &什克洛夫ShklowBarony{}

func init() {
	f := BShklow什克洛夫.(*什克洛夫ShklowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shklow",
		TitleName: "什克洛夫",
		TitleCode: "b_shklow",
	}
}
