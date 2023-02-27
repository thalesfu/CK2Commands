package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PisaKingdom interface {
    feud.Kingdom
}

type 比萨PisaKingdom struct {
	feud.BaseKingdom
}

var KPisa比萨 PisaKingdom = &比萨PisaKingdom{}

func init() {
	f := KPisa比萨.(*比萨PisaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "pisa",
		TitleName: "比萨",
		TitleCode: "k_pisa",
		Dukes:  map[string]feud.Duke{},
	}

}
