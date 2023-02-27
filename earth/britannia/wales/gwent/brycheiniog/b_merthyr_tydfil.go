package brycheiniog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅瑟蒂德菲尔Merthyr_tydfilBarony struct {
	feud.BaseBarony
}

var BMerthyr_tydfil梅瑟蒂德菲尔 feud.Barony = &梅瑟蒂德菲尔Merthyr_tydfilBarony{}

func init() {
    f := BMerthyr_tydfil梅瑟蒂德菲尔.(*梅瑟蒂德菲尔Merthyr_tydfilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merthyr_tydfil",
		TitleName: "梅瑟蒂德菲尔",
		TitleCode: "b_merthyr_tydfil",
	}
}
