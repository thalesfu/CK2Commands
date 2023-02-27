package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达拉卜DarabBarony struct {
	feud.BaseBarony
}

var BDarab达拉卜 feud.Barony = &达拉卜DarabBarony{}

func init() {
    f := BDarab达拉卜.(*达拉卜DarabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darab",
		TitleName: "达拉卜",
		TitleCode: "b_darab",
	}
}
