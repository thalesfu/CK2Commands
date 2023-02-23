package swabia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SwabiaKingdom interface {
	feud.Kingdom
}

type 施瓦本SwabiaKingdom struct {
	feud.BaseKingdom
}

var KSwabia施瓦本 SwabiaKingdom = &施瓦本SwabiaKingdom{}

func init() {
	f := KSwabia施瓦本.(*施瓦本SwabiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "swabia",
		TitleName: "施瓦本",
		TitleCode: "k_swabia",
		Dukes:     map[string]feud.Duke{},
	}

}
