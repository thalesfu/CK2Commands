package guzgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安狄枯AndkhudBarony struct {
	feud.BaseBarony
}

var BAndkhud安狄枯 feud.Barony = &安狄枯AndkhudBarony{}

func init() {
    f := BAndkhud安狄枯.(*安狄枯AndkhudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andkhud",
		TitleName: "安狄枯",
		TitleCode: "b_andkhud",
	}
}
