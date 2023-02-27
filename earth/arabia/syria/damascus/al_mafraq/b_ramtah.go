package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖姆萨RamtahBarony struct {
	feud.BaseBarony
}

var BRamtah赖姆萨 feud.Barony = &赖姆萨RamtahBarony{}

func init() {
    f := BRamtah赖姆萨.(*赖姆萨RamtahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramtah",
		TitleName: "赖姆萨",
		TitleCode: "b_ramtah",
	}
}
