package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法卡尼亚斯FacaniasBarony struct {
	feud.BaseBarony
}

var BFacanias法卡尼亚斯 feud.Barony = &法卡尼亚斯FacaniasBarony{}

func init() {
    f := BFacanias法卡尼亚斯.(*法卡尼亚斯FacaniasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "facanias",
		TitleName: "法卡尼亚斯",
		TitleCode: "b_facanias",
	}
}
