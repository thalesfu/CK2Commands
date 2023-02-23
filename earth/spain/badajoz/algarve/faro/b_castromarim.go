package faro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马林堡CastromarimBarony struct {
	feud.BaseBarony
}

var BCastromarim马林堡 feud.Barony = &马林堡CastromarimBarony{}

func init() {
	f := BCastromarim马林堡.(*马林堡CastromarimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castromarim",
		TitleName: "马林堡",
		TitleCode: "b_castromarim",
	}
}
