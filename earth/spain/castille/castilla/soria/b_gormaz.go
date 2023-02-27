package soria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔马斯GormazBarony struct {
	feud.BaseBarony
}

var BGormaz戈尔马斯 feud.Barony = &戈尔马斯GormazBarony{}

func init() {
    f := BGormaz戈尔马斯.(*戈尔马斯GormazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gormaz",
		TitleName: "戈尔马斯",
		TitleCode: "b_gormaz",
	}
}
