package tui

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏普拉SupraBarony struct {
	feud.BaseBarony
}

var BSupra苏普拉 feud.Barony = &苏普拉SupraBarony{}

func init() {
    f := BSupra苏普拉.(*苏普拉SupraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "supra",
		TitleName: "苏普拉",
		TitleCode: "b_supra",
	}
}
