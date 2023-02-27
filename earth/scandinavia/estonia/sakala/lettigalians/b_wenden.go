package lettigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 文登WendenBarony struct {
	feud.BaseBarony
}

var BWenden文登 feud.Barony = &文登WendenBarony{}

func init() {
    f := BWenden文登.(*文登WendenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wenden",
		TitleName: "文登",
		TitleCode: "b_wenden",
	}
}
