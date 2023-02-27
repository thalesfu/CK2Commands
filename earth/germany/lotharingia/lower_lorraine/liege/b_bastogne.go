package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯托涅BastogneBarony struct {
	feud.BaseBarony
}

var BBastogne巴斯托涅 feud.Barony = &巴斯托涅BastogneBarony{}

func init() {
    f := BBastogne巴斯托涅.(*巴斯托涅BastogneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bastogne",
		TitleName: "巴斯托涅",
		TitleCode: "b_bastogne",
	}
}
