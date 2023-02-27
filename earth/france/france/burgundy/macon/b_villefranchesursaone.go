package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索恩河畔自由城VillefranchesursaoneBarony struct {
	feud.BaseBarony
}

var BVillefranchesursaone索恩河畔自由城 feud.Barony = &索恩河畔自由城VillefranchesursaoneBarony{}

func init() {
    f := BVillefranchesursaone索恩河畔自由城.(*索恩河畔自由城VillefranchesursaoneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villefranchesursaone",
		TitleName: "索恩河畔自由城",
		TitleCode: "b_villefranchesursaone",
	}
}
