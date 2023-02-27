package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯比KirkbyBarony struct {
	feud.BaseBarony
}

var BKirkby柯比 feud.Barony = &柯比KirkbyBarony{}

func init() {
    f := BKirkby柯比.(*柯比KirkbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirkby",
		TitleName: "柯比",
		TitleCode: "b_kirkby",
	}
}
