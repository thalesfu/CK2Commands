package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡奏CadzowBarony struct {
	feud.BaseBarony
}

var BCadzow卡奏 feud.Barony = &卡奏CadzowBarony{}

func init() {
	f := BCadzow卡奏.(*卡奏CadzowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cadzow",
		TitleName: "卡奏",
		TitleCode: "b_cadzow",
	}
}
