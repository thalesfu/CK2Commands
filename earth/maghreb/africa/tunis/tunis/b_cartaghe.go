package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦太基CartagheBarony struct {
	feud.BaseBarony
}

var BCartaghe迦太基 feud.Barony = &迦太基CartagheBarony{}

func init() {
	f := BCartaghe迦太基.(*迦太基CartagheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cartaghe",
		TitleName: "迦太基",
		TitleCode: "b_cartaghe",
	}
}
