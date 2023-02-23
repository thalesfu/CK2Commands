package novogrodek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊维耶IwieBarony struct {
	feud.BaseBarony
}

var BIwie伊维耶 feud.Barony = &伊维耶IwieBarony{}

func init() {
	f := BIwie伊维耶.(*伊维耶IwieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iwie",
		TitleName: "伊维耶",
		TitleCode: "b_iwie",
	}
}
