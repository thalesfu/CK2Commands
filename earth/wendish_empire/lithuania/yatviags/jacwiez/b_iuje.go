package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊维耶IujeBarony struct {
	feud.BaseBarony
}

var BIuje伊维耶 feud.Barony = &伊维耶IujeBarony{}

func init() {
    f := BIuje伊维耶.(*伊维耶IujeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iuje",
		TitleName: "伊维耶",
		TitleCode: "b_iuje",
	}
}
