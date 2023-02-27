package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥罗斯拉维OroslavjeBarony struct {
	feud.BaseBarony
}

var BOroslavje奥罗斯拉维 feud.Barony = &奥罗斯拉维OroslavjeBarony{}

func init() {
    f := BOroslavje奥罗斯拉维.(*奥罗斯拉维OroslavjeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oroslavje",
		TitleName: "奥罗斯拉维",
		TitleCode: "b_oroslavje",
	}
}
