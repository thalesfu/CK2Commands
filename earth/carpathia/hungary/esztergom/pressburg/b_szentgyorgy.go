package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣捷尔吉SzentgyorgyBarony struct {
	feud.BaseBarony
}

var BSzentgyorgy圣捷尔吉 feud.Barony = &圣捷尔吉SzentgyorgyBarony{}

func init() {
    f := BSzentgyorgy圣捷尔吉.(*圣捷尔吉SzentgyorgyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szentgyorgy",
		TitleName: "圣捷尔吉",
		TitleCode: "b_szentgyorgy",
	}
}
