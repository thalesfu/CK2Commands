package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢普希圣哲尔吉SepsiszentgyorgyBarony struct {
	feud.BaseBarony
}

var BSepsiszentgyorgy谢普希圣哲尔吉 feud.Barony = &谢普希圣哲尔吉SepsiszentgyorgyBarony{}

func init() {
	f := BSepsiszentgyorgy谢普希圣哲尔吉.(*谢普希圣哲尔吉SepsiszentgyorgyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sepsiszentgyorgy",
		TitleName: "谢普希圣哲尔吉",
		TitleCode: "b_sepsiszentgyorgy",
	}
}
