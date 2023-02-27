package canarias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉帕尔马LapalmaBarony struct {
	feud.BaseBarony
}

var BLapalma拉帕尔马 feud.Barony = &拉帕尔马LapalmaBarony{}

func init() {
    f := BLapalma拉帕尔马.(*拉帕尔马LapalmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lapalma",
		TitleName: "拉帕尔马",
		TitleCode: "b_lapalma",
	}
}
