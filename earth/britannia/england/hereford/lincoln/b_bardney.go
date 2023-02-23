package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴德尼BardneyBarony struct {
	feud.BaseBarony
}

var BBardney巴德尼 feud.Barony = &巴德尼BardneyBarony{}

func init() {
	f := BBardney巴德尼.(*巴德尼BardneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bardney",
		TitleName: "巴德尼",
		TitleCode: "b_bardney",
	}
}
