package mahoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗提RathBarony struct {
	feud.BaseBarony
}

var BRath罗提 feud.Barony = &罗提RathBarony{}

func init() {
    f := BRath罗提.(*罗提RathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rath",
		TitleName: "罗提",
		TitleCode: "b_rath",
	}
}
