package swabia_otto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Swabia_ottoKingdom interface {
    feud.Kingdom
}

type 施瓦本Swabia_ottoKingdom struct {
	feud.BaseKingdom
}

var KSwabia_otto施瓦本 Swabia_ottoKingdom = &施瓦本Swabia_ottoKingdom{}

func init() {
	f := KSwabia_otto施瓦本.(*施瓦本Swabia_ottoKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "swabia_otto",
		TitleName: "施瓦本",
		TitleCode: "k_swabia_otto",
		Dukes:  map[string]feud.Duke{},
	}

}
