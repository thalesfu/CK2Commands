package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔帕索斯KarpathosBarony struct {
	feud.BaseBarony
}

var BKarpathos卡尔帕索斯 feud.Barony = &卡尔帕索斯KarpathosBarony{}

func init() {
    f := BKarpathos卡尔帕索斯.(*卡尔帕索斯KarpathosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karpathos",
		TitleName: "卡尔帕索斯",
		TitleCode: "b_karpathos",
	}
}
