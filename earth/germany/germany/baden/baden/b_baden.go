package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴登BadenBarony struct {
	feud.BaseBarony
}

var BBaden巴登 feud.Barony = &巴登BadenBarony{}

func init() {
    f := BBaden巴登.(*巴登BadenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baden",
		TitleName: "巴登",
		TitleCode: "b_baden",
	}
}
