package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尔斯FarsiBarony struct {
	feud.BaseBarony
}

var BFarsi法尔斯 feud.Barony = &法尔斯FarsiBarony{}

func init() {
	f := BFarsi法尔斯.(*法尔斯FarsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farsi",
		TitleName: "法尔斯",
		TitleCode: "b_farsi",
	}
}
