package korsun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹纳米安卡ZnamiankaBarony struct {
	feud.BaseBarony
}

var BZnamianka兹纳米安卡 feud.Barony = &兹纳米安卡ZnamiankaBarony{}

func init() {
    f := BZnamianka兹纳米安卡.(*兹纳米安卡ZnamiankaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "znamianka",
		TitleName: "兹纳米安卡",
		TitleCode: "b_znamianka",
	}
}
