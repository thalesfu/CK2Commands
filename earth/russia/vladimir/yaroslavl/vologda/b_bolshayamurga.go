package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大穆尔加BolshayamurgaBarony struct {
	feud.BaseBarony
}

var BBolshayamurga大穆尔加 feud.Barony = &大穆尔加BolshayamurgaBarony{}

func init() {
	f := BBolshayamurga大穆尔加.(*大穆尔加BolshayamurgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolshayamurga",
		TitleName: "大穆尔加",
		TitleCode: "b_bolshayamurga",
	}
}
