package mountain_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅列克斯MelekessBarony struct {
	feud.BaseBarony
}

var BMelekess梅列克斯 feud.Barony = &梅列克斯MelekessBarony{}

func init() {
    f := BMelekess梅列克斯.(*梅列克斯MelekessBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melekess",
		TitleName: "梅列克斯",
		TitleCode: "b_melekess",
	}
}
