package tirgoviste

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布泽乌BuzauBarony struct {
	feud.BaseBarony
}

var BBuzau布泽乌 feud.Barony = &布泽乌BuzauBarony{}

func init() {
    f := BBuzau布泽乌.(*布泽乌BuzauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buzau",
		TitleName: "布泽乌",
		TitleCode: "b_buzau",
	}
}
