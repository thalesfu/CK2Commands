package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔登克龙GoldenkronBarony struct {
	feud.BaseBarony
}

var BGoldenkron戈尔登克龙 feud.Barony = &戈尔登克龙GoldenkronBarony{}

func init() {
    f := BGoldenkron戈尔登克龙.(*戈尔登克龙GoldenkronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goldenkron",
		TitleName: "戈尔登克龙",
		TitleCode: "b_goldenkron",
	}
}
