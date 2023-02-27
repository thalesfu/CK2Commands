package zemigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多贝DobeBarony struct {
	feud.BaseBarony
}

var BDobe多贝 feud.Barony = &多贝DobeBarony{}

func init() {
    f := BDobe多贝.(*多贝DobeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobe",
		TitleName: "多贝",
		TitleCode: "b_dobe",
	}
}
