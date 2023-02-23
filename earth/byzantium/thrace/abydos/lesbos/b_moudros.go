package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆兹罗斯MoudrosBarony struct {
	feud.BaseBarony
}

var BMoudros穆兹罗斯 feud.Barony = &穆兹罗斯MoudrosBarony{}

func init() {
	f := BMoudros穆兹罗斯.(*穆兹罗斯MoudrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moudros",
		TitleName: "穆兹罗斯",
		TitleCode: "b_moudros",
	}
}
