package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔罗斯MelroseBarony struct {
	feud.BaseBarony
}

var BMelrose梅尔罗斯 feud.Barony = &梅尔罗斯MelroseBarony{}

func init() {
    f := BMelrose梅尔罗斯.(*梅尔罗斯MelroseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melrose",
		TitleName: "梅尔罗斯",
		TitleCode: "b_melrose",
	}
}
