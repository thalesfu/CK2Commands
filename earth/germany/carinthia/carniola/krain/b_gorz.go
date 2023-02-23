package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格尔茨GorzBarony struct {
	feud.BaseBarony
}

var BGorz格尔茨 feud.Barony = &格尔茨GorzBarony{}

func init() {
	f := BGorz格尔茨.(*格尔茨GorzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorz",
		TitleName: "格尔茨",
		TitleCode: "b_gorz",
	}
}
