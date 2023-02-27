package velsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希洛夫斯卡亚ShilovskayaBarony struct {
	feud.BaseBarony
}

var BShilovskaya希洛夫斯卡亚 feud.Barony = &希洛夫斯卡亚ShilovskayaBarony{}

func init() {
    f := BShilovskaya希洛夫斯卡亚.(*希洛夫斯卡亚ShilovskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shilovskaya",
		TitleName: "希洛夫斯卡亚",
		TitleCode: "b_shilovskaya",
	}
}
