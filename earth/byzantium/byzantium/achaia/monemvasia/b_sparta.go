package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯巴达SpartaBarony struct {
	feud.BaseBarony
}

var BSparta斯巴达 feud.Barony = &斯巴达SpartaBarony{}

func init() {
    f := BSparta斯巴达.(*斯巴达SpartaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sparta",
		TitleName: "斯巴达",
		TitleCode: "b_sparta",
	}
}
