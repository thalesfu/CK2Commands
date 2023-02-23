package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦乌泰特OuaouteitBarony struct {
	feud.BaseBarony
}

var BOuaouteit瓦乌泰特 feud.Barony = &瓦乌泰特OuaouteitBarony{}

func init() {
	f := BOuaouteit瓦乌泰特.(*瓦乌泰特OuaouteitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouaouteit",
		TitleName: "瓦乌泰特",
		TitleCode: "b_ouaouteit",
	}
}
