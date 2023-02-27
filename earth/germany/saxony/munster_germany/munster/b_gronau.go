package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格罗瑙GronauBarony struct {
	feud.BaseBarony
}

var BGronau格罗瑙 feud.Barony = &格罗瑙GronauBarony{}

func init() {
    f := BGronau格罗瑙.(*格罗瑙GronauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gronau",
		TitleName: "格罗瑙",
		TitleCode: "b_gronau",
	}
}
