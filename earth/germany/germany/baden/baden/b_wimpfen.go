package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温普芬WimpfenBarony struct {
	feud.BaseBarony
}

var BWimpfen温普芬 feud.Barony = &温普芬WimpfenBarony{}

func init() {
    f := BWimpfen温普芬.(*温普芬WimpfenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wimpfen",
		TitleName: "温普芬",
		TitleCode: "b_wimpfen",
	}
}
