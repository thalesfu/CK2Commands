package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰内DjenneBarony struct {
	feud.BaseBarony
}

var BDjenne杰内 feud.Barony = &杰内DjenneBarony{}

func init() {
    f := BDjenne杰内.(*杰内DjenneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djenne",
		TitleName: "杰内",
		TitleCode: "b_djenne",
	}
}
