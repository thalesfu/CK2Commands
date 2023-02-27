package medantaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迷檀多迦MedantakaBarony struct {
	feud.BaseBarony
}

var BMedantaka迷檀多迦 feud.Barony = &迷檀多迦MedantakaBarony{}

func init() {
    f := BMedantaka迷檀多迦.(*迷檀多迦MedantakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medantaka",
		TitleName: "迷檀多迦",
		TitleCode: "b_medantaka",
	}
}
