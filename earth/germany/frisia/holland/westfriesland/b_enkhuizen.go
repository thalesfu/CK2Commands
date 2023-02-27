package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩克黑曾EnkhuizenBarony struct {
	feud.BaseBarony
}

var BEnkhuizen恩克黑曾 feud.Barony = &恩克黑曾EnkhuizenBarony{}

func init() {
    f := BEnkhuizen恩克黑曾.(*恩克黑曾EnkhuizenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "enkhuizen",
		TitleName: "恩克黑曾",
		TitleCode: "b_enkhuizen",
	}
}
