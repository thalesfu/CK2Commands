package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜吉BayjiBarony struct {
	feud.BaseBarony
}

var BBayji拜吉 feud.Barony = &拜吉BayjiBarony{}

func init() {
    f := BBayji拜吉.(*拜吉BayjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayji",
		TitleName: "拜吉",
		TitleCode: "b_bayji",
	}
}
