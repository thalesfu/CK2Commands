package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利尔LierBarony struct {
	feud.BaseBarony
}

var BLier利尔 feud.Barony = &利尔LierBarony{}

func init() {
	f := BLier利尔.(*利尔LierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lier",
		TitleName: "利尔",
		TitleCode: "b_lier",
	}
}
