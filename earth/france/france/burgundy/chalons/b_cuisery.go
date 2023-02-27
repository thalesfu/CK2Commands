package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奎瑟里CuiseryBarony struct {
	feud.BaseBarony
}

var BCuisery奎瑟里 feud.Barony = &奎瑟里CuiseryBarony{}

func init() {
    f := BCuisery奎瑟里.(*奎瑟里CuiseryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cuisery",
		TitleName: "奎瑟里",
		TitleCode: "b_cuisery",
	}
}
