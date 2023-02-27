package vychegda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃利季诺Vol_dinoBarony struct {
	feud.BaseBarony
}

var BVol_dino沃利季诺 feud.Barony = &沃利季诺Vol_dinoBarony{}

func init() {
    f := BVol_dino沃利季诺.(*沃利季诺Vol_dinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vol_dino",
		TitleName: "沃利季诺",
		TitleCode: "b_vol_dino",
	}
}
