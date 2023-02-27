package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴喀法BaqofahBarony struct {
	feud.BaseBarony
}

var BBaqofah巴喀法 feud.Barony = &巴喀法BaqofahBarony{}

func init() {
    f := BBaqofah巴喀法.(*巴喀法BaqofahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baqofah",
		TitleName: "巴喀法",
		TitleCode: "b_baqofah",
	}
}
