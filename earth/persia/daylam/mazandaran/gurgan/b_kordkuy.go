package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔德库伊KordkuyBarony struct {
	feud.BaseBarony
}

var BKordkuy科尔德库伊 feud.Barony = &科尔德库伊KordkuyBarony{}

func init() {
    f := BKordkuy科尔德库伊.(*科尔德库伊KordkuyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kordkuy",
		TitleName: "科尔德库伊",
		TitleCode: "b_kordkuy",
	}
}
