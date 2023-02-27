package lakomelza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德西DessyeBarony struct {
	feud.BaseBarony
}

var BDessye德西 feud.Barony = &德西DessyeBarony{}

func init() {
    f := BDessye德西.(*德西DessyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dessye",
		TitleName: "德西",
		TitleCode: "b_dessye",
	}
}
