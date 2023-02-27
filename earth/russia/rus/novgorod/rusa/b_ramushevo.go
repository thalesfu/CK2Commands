package rusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉穆舍沃RamushevoBarony struct {
	feud.BaseBarony
}

var BRamushevo拉穆舍沃 feud.Barony = &拉穆舍沃RamushevoBarony{}

func init() {
    f := BRamushevo拉穆舍沃.(*拉穆舍沃RamushevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramushevo",
		TitleName: "拉穆舍沃",
		TitleCode: "b_ramushevo",
	}
}
