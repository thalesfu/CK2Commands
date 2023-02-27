package yuni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 花川村HuachuantsunBarony struct {
	feud.BaseBarony
}

var BHuachuantsun花川村 feud.Barony = &花川村HuachuantsunBarony{}

func init() {
    f := BHuachuantsun花川村.(*花川村HuachuantsunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huachuantsun",
		TitleName: "花川村",
		TitleCode: "b_huachuantsun",
	}
}
