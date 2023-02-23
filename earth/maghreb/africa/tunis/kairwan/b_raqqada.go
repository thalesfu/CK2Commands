package kairwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷加达RaqqadaBarony struct {
	feud.BaseBarony
}

var BRaqqada雷加达 feud.Barony = &雷加达RaqqadaBarony{}

func init() {
	f := BRaqqada雷加达.(*雷加达RaqqadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raqqada",
		TitleName: "雷加达",
		TitleCode: "b_raqqada",
	}
}
