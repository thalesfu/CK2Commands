package kinda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂耶什塔德TjarstadBarony struct {
	feud.BaseBarony
}

var BTjarstad蒂耶什塔德 feud.Barony = &蒂耶什塔德TjarstadBarony{}

func init() {
	f := BTjarstad蒂耶什塔德.(*蒂耶什塔德TjarstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tjarstad",
		TitleName: "蒂耶什塔德",
		TitleCode: "b_tjarstad",
	}
}
