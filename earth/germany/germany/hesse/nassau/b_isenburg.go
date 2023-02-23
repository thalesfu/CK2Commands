package nassau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊森堡IsenburgBarony struct {
	feud.BaseBarony
}

var BIsenburg伊森堡 feud.Barony = &伊森堡IsenburgBarony{}

func init() {
	f := BIsenburg伊森堡.(*伊森堡IsenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isenburg",
		TitleName: "伊森堡",
		TitleCode: "b_isenburg",
	}
}
