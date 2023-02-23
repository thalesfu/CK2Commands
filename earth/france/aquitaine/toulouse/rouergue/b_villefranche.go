package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 自由城VillefrancheBarony struct {
	feud.BaseBarony
}

var BVillefranche自由城 feud.Barony = &自由城VillefrancheBarony{}

func init() {
	f := BVillefranche自由城.(*自由城VillefrancheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villefranche",
		TitleName: "自由城",
		TitleCode: "b_villefranche",
	}
}
