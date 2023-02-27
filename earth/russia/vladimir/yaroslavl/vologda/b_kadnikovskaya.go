package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡德尼科夫斯卡亚KadnikovskayaBarony struct {
	feud.BaseBarony
}

var BKadnikovskaya卡德尼科夫斯卡亚 feud.Barony = &卡德尼科夫斯卡亚KadnikovskayaBarony{}

func init() {
    f := BKadnikovskaya卡德尼科夫斯卡亚.(*卡德尼科夫斯卡亚KadnikovskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadnikovskaya",
		TitleName: "卡德尼科夫斯卡亚",
		TitleCode: "b_kadnikovskaya",
	}
}
