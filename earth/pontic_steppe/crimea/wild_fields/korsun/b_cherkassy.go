package korsun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔卡瑟CherkassyBarony struct {
	feud.BaseBarony
}

var BCherkassy切尔卡瑟 feud.Barony = &切尔卡瑟CherkassyBarony{}

func init() {
    f := BCherkassy切尔卡瑟.(*切尔卡瑟CherkassyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherkassy",
		TitleName: "切尔卡瑟",
		TitleCode: "b_cherkassy",
	}
}
