package retz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣布雷万Saint_brevinBarony struct {
	feud.BaseBarony
}

var BSaint_brevin圣布雷万 feud.Barony = &圣布雷万Saint_brevinBarony{}

func init() {
    f := BSaint_brevin圣布雷万.(*圣布雷万Saint_brevinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saint_brevin",
		TitleName: "圣布雷万",
		TitleCode: "b_saint_brevin",
	}
}
