package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗利ForliBarony struct {
	feud.BaseBarony
}

var BForli弗利 feud.Barony = &弗利ForliBarony{}

func init() {
    f := BForli弗利.(*弗利ForliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forli",
		TitleName: "弗利",
		TitleCode: "b_forli",
	}
}
