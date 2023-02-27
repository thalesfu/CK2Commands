package lhasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绕木切RamocheBarony struct {
	feud.BaseBarony
}

var BRamoche绕木切 feud.Barony = &绕木切RamocheBarony{}

func init() {
    f := BRamoche绕木切.(*绕木切RamocheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramoche",
		TitleName: "绕木切",
		TitleCode: "b_ramoche",
	}
}
