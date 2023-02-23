package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松博特海伊SzombathelyBarony struct {
	feud.BaseBarony
}

var BSzombathely松博特海伊 feud.Barony = &松博特海伊SzombathelyBarony{}

func init() {
	f := BSzombathely松博特海伊.(*松博特海伊SzombathelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szombathely",
		TitleName: "松博特海伊",
		TitleCode: "b_szombathely",
	}
}
