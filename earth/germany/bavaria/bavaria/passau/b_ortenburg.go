package passau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔滕堡OrtenburgBarony struct {
	feud.BaseBarony
}

var BOrtenburg奥尔滕堡 feud.Barony = &奥尔滕堡OrtenburgBarony{}

func init() {
	f := BOrtenburg奥尔滕堡.(*奥尔滕堡OrtenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ortenburg",
		TitleName: "奥尔滕堡",
		TitleCode: "b_ortenburg",
	}
}
