package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉姆斯伯里RamsburyBarony struct {
	feud.BaseBarony
}

var BRamsbury拉姆斯伯里 feud.Barony = &拉姆斯伯里RamsburyBarony{}

func init() {
    f := BRamsbury拉姆斯伯里.(*拉姆斯伯里RamsburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramsbury",
		TitleName: "拉姆斯伯里",
		TitleCode: "b_ramsbury",
	}
}
