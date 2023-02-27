package kano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加马贾GamajaBarony struct {
	feud.BaseBarony
}

var BGamaja加马贾 feud.Barony = &加马贾GamajaBarony{}

func init() {
    f := BGamaja加马贾.(*加马贾GamajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gamaja",
		TitleName: "加马贾",
		TitleCode: "b_gamaja",
	}
}
