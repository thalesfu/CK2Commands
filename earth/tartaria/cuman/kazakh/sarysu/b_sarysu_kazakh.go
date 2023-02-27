package sarysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷苏Sarysu_kazakhBarony struct {
	feud.BaseBarony
}

var BSarysu_kazakh萨雷苏 feud.Barony = &萨雷苏Sarysu_kazakhBarony{}

func init() {
    f := BSarysu_kazakh萨雷苏.(*萨雷苏Sarysu_kazakhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarysu_kazakh",
		TitleName: "萨雷苏",
		TitleCode: "b_sarysu_kazakh",
	}
}
