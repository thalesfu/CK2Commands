package yumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黄闸湾HuangzhawanBarony struct {
	feud.BaseBarony
}

var BHuangzhawan黄闸湾 feud.Barony = &黄闸湾HuangzhawanBarony{}

func init() {
	f := BHuangzhawan黄闸湾.(*黄闸湾HuangzhawanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huangzhawan",
		TitleName: "黄闸湾",
		TitleCode: "b_huangzhawan",
	}
}
