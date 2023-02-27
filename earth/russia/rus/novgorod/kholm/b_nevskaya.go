package kholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅夫斯卡亚NevskayaBarony struct {
	feud.BaseBarony
}

var BNevskaya涅夫斯卡亚 feud.Barony = &涅夫斯卡亚NevskayaBarony{}

func init() {
    f := BNevskaya涅夫斯卡亚.(*涅夫斯卡亚NevskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nevskaya",
		TitleName: "涅夫斯卡亚",
		TitleCode: "b_nevskaya",
	}
}
