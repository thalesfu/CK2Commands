package ui_fiachrach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴林托伯BallintubberBarony struct {
	feud.BaseBarony
}

var BBallintubber巴林托伯 feud.Barony = &巴林托伯BallintubberBarony{}

func init() {
    f := BBallintubber巴林托伯.(*巴林托伯BallintubberBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ballintubber",
		TitleName: "巴林托伯",
		TitleCode: "b_ballintubber",
	}
}
