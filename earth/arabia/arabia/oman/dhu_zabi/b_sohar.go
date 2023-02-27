package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏哈尔SoharBarony struct {
	feud.BaseBarony
}

var BSohar苏哈尔 feud.Barony = &苏哈尔SoharBarony{}

func init() {
    f := BSohar苏哈尔.(*苏哈尔SoharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sohar",
		TitleName: "苏哈尔",
		TitleCode: "b_sohar",
	}
}
