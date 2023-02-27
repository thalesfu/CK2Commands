package naples

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NaplesKingdom interface {
    feud.Kingdom
}

type 那不勒斯NaplesKingdom struct {
	feud.BaseKingdom
}

var KNaples那不勒斯 NaplesKingdom = &那不勒斯NaplesKingdom{}

func init() {
	f := KNaples那不勒斯.(*那不勒斯NaplesKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "naples",
		TitleName: "那不勒斯",
		TitleCode: "k_naples",
		Dukes:  map[string]feud.Duke{},
	}

}
