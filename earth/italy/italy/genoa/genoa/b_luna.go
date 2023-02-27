package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢纳LunaBarony struct {
	feud.BaseBarony
}

var BLuna卢纳 feud.Barony = &卢纳LunaBarony{}

func init() {
    f := BLuna卢纳.(*卢纳LunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luna",
		TitleName: "卢纳",
		TitleCode: "b_luna",
	}
}
