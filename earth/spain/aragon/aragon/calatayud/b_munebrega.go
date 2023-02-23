package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆内夫雷加MunebregaBarony struct {
	feud.BaseBarony
}

var BMunebrega穆内夫雷加 feud.Barony = &穆内夫雷加MunebregaBarony{}

func init() {
	f := BMunebrega穆内夫雷加.(*穆内夫雷加MunebregaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "munebrega",
		TitleName: "穆内夫雷加",
		TitleCode: "b_munebrega",
	}
}
